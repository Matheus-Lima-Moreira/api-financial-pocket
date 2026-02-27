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

func (s *Service) Details(ctx context.Context, id uint) (*GroupPermissionEntity, *shared_errors.AppError) {
	groupPermission, err := s.groupPermissionRepository.Details(ctx, id)
	if err != nil {
		return nil, shared_errors.NewBadRequest(err.Error())
	}
	return groupPermission, nil
}

func (s *Service) Create(ctx context.Context, groupPermission *GroupPermissionEntity) *shared_errors.AppError {
	err := s.groupPermissionRepository.Create(ctx, groupPermission)
	if err != nil {
		return shared_errors.NewBadRequest(err.Error())
	}
	return nil
}

func (s *Service) Update(ctx context.Context, groupPermission *GroupPermissionEntity) *shared_errors.AppError {
	err := s.groupPermissionRepository.Update(ctx, groupPermission)
	if err != nil {
		return shared_errors.NewBadRequest(err.Error())
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id uint) *shared_errors.AppError {
	err := s.groupPermissionRepository.Delete(ctx, id)
	if err != nil {
		return shared_errors.NewBadRequest(err.Error())
	}
	return nil
}
