package organizations

import (
	"context"
	"errors"
	"time"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/consts"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

const mysqlErrDuplicateEntry = 1062

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func isDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == mysqlErrDuplicateEntry
	}
	return false
}

func (r *GormRepository) Create(ctx context.Context, organization *OrganizationEntity) *shared_errors.AppError {
	model := toModel(organization)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		if isDuplicateKeyError(err) {
			return shared_errors.NewConflict("error.organization_already_exists", "organization")
		}
		return shared_errors.NewBadRequest(err.Error())
	}

	*organization = *toDomain(model)

	return nil
}

func (r *GormRepository) List(ctx context.Context, page int) ([]OrganizationEntity, *dtos.PaginationDTO, *shared_errors.AppError) {
	var models []OrganizationSchema

	limit := consts.PaginationDefaultLimit

	err := r.db.WithContext(ctx).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&models).Error

	if err != nil {
		return nil, nil, shared_errors.NewBadRequest(err.Error())
	}

	domains := make([]OrganizationEntity, len(models))
	for i, model := range models {
		domains[i] = *toDomain(&model)
	}

	var total int64 = 0

	err = r.db.WithContext(ctx).
		Model(&OrganizationSchema{}).
		Count(&total).Error

	if err != nil {
		return nil, nil, shared_errors.NewBadRequest(err.Error())
	}

	next := 0
	previous := 0
	if page*limit < int(total) {
		next = page + 1
	}
	if page > 1 {
		previous = page - 1
	}

	pagination := &dtos.PaginationDTO{
		Page:     page,
		Limit:    limit,
		Total:    int(total),
		Next:     next,
		Previous: previous,
	}

	return domains, pagination, nil
}

func (r *GormRepository) GetById(ctx context.Context, organizationID string) (*OrganizationEntity, *shared_errors.AppError) {
	var model OrganizationSchema

	err := r.db.WithContext(ctx).
		Where("id = ?", organizationID).
		First(&model).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, shared_errors.NewNotFound("organization")
		}
		return nil, shared_errors.NewBadRequest(err.Error())
	}

	return toDomain(&model), nil
}

func (r *GormRepository) Update(ctx context.Context, organization *OrganizationEntity) *shared_errors.AppError {
	err := r.db.WithContext(ctx).
		Model(&OrganizationSchema{}).
		Where("id = ?", organization.ID).
		Update("name", organization.Name).
		Update("cellphone", organization.Cellphone).
		Update("logo", organization.Logo).
		Update("updated_at", time.Now()).Error

	if err != nil {
		return shared_errors.NewBadRequest(err.Error())
	}

	return nil
}

func (r *GormRepository) Delete(ctx context.Context, organizationID string) *shared_errors.AppError {
	err := r.db.WithContext(ctx).
		Model(&OrganizationSchema{}).
		Where("id = ?", organizationID).
		Delete(&OrganizationSchema{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return shared_errors.NewNotFound("organization")
		}
		return shared_errors.NewBadRequest(err.Error())
	}

	return nil
}
