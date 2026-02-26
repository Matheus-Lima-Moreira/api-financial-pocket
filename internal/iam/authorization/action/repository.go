package action

import (
	"context"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

type Repository interface {
	List(ctx context.Context, page int) ([]ActionEntity, *dtos.PaginationDTO, *shared_errors.AppError)
}
