package action

import (
	"context"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/consts"
	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &GormRepository{db: db}
}

func (r *GormRepository) List(ctx context.Context, page int) ([]ActionEntity, *dtos.PaginationDTO, *shared_errors.AppError) {
	var models []ActionSchema

	limit := consts.PaginationDefaultLimit

	err := r.db.WithContext(ctx).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&models).Error

	if err != nil {
		return nil, nil, shared_errors.NewBadRequest(err.Error())
	}

	domains := make([]ActionEntity, len(models))
	for i, model := range models {
		domains[i] = *toDomain(&model)
	}

	var total int64 = 0

	err = r.db.WithContext(ctx).
		Model(&ActionSchema{}).
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
