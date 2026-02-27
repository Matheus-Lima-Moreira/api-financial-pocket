package organizations

import (
	"context"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

type Repository interface {
	Create(ctx context.Context, organization *OrganizationEntity) *shared_errors.AppError
	List(ctx context.Context, page int) ([]OrganizationEntity, *dtos.PaginationDTO, *shared_errors.AppError)
	GetById(ctx context.Context, id string) (*OrganizationEntity, *shared_errors.AppError)
	Update(ctx context.Context, organization *OrganizationEntity) *shared_errors.AppError
	Delete(ctx context.Context, id string) *shared_errors.AppError
}
