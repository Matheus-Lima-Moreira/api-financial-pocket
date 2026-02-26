package action

import (
	"context"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

type Service struct {
	actionRepository Repository
}

func NewService(actionRepository Repository) *Service {
	return &Service{
		actionRepository: actionRepository,
	}
}

func (s *Service) List(ctx context.Context, page int) ([]ActionEntity, *dtos.PaginationDTO, *shared_errors.AppError) {
	actions, pagination, err := s.actionRepository.List(ctx, page)
	if err != nil {
		return nil, nil, shared_errors.NewBadRequest(err.Error())
	}
	return actions, pagination, nil
}
