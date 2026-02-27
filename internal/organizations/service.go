package organizations

import (
	"context"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

type Service struct {
	organizationRepository Repository
}

func NewService(organizationRepository Repository) *Service {
	return &Service{
		organizationRepository: organizationRepository,
	}
}

func (s *Service) List(ctx context.Context, page int) ([]OrganizationEntity, *dtos.PaginationDTO, *shared_errors.AppError) {
	organizations, pagination, err := s.organizationRepository.List(ctx, page)
	if err != nil {
		return nil, nil, shared_errors.NewBadRequest(err.Error())
	}
	return organizations, pagination, nil
}

func (s *Service) Details(ctx context.Context, id string) (*OrganizationEntity, *shared_errors.AppError) {
	organization, err := s.organizationRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func (s *Service) Create(ctx context.Context, organization *OrganizationEntity) *shared_errors.AppError {
	err := s.organizationRepository.Create(ctx, organization)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, organization *OrganizationEntity) *shared_errors.AppError {
	err := s.organizationRepository.Update(ctx, organization)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Delete(ctx context.Context, id string) *shared_errors.AppError {
	err := s.organizationRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
