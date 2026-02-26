package group_permission

import (
	"context"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

type Service struct {
	groupPermissionRepository Repository
}

func NewService(groupPermissionRepository Repository) *Service {
	return &Service{
		groupPermissionRepository: groupPermissionRepository,
	}
}

func (s *Service) List(ctx context.Context, page int) ([]GroupPermissionEntity, *dtos.PaginationDTO, *shared_errors.AppError) {
	groupPermissions, pagination, err := s.groupPermissionRepository.List(ctx, page)
	if err != nil {
		return nil, nil, shared_errors.NewBadRequest(err.Error())
	}
	return groupPermissions, pagination, nil
}
