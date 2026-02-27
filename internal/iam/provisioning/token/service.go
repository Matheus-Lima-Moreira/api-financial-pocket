package token

import (
	"context"
	"time"

	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

type Service struct {
	tokenRepository Repository
}

func NewService(tokenRepository Repository) *Service {
	return &Service{
		tokenRepository: tokenRepository,
	}
}

func (s *Service) GenerateToken(ctx context.Context, resource TokenResource, referenceID string, metadata map[string]any) (*TokenEntity, *shared_errors.AppError) {
	token, err := GenerateToken()
	if err != nil {
		return nil, err
	}

	tokenEntity := NewTokenEntity(resource, token, referenceID, metadata)

	if err := s.tokenRepository.Create(ctx, tokenEntity); err != nil {
		return nil, err
	}

	return tokenEntity, nil
}

func (s *Service) VerifyToken(ctx context.Context, token string, resource *TokenResource) (bool, *shared_errors.AppError) {
	tokenEntity, err := s.tokenRepository.FindByToken(ctx, token)
	if err != nil {
		return false, shared_errors.NewUnauthorized("error.invalid_token")
	}

	if tokenEntity.Status != TokenStatusActive {
		return false, shared_errors.NewUnauthorized("error.invalid_token")
	}

	if tokenEntity.ExpiresAt.Before(time.Now()) {
		return false, shared_errors.NewUnauthorized("error.expired_token")
	}

	if resource != nil && tokenEntity.Resource != *resource {
		return false, shared_errors.NewUnauthorized("error.invalid_token")
	}

	return true, nil
}

func (s *Service) GetByToken(ctx context.Context, token string) (*TokenEntity, *shared_errors.AppError) {
	tokenEntity, err := s.tokenRepository.FindByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return tokenEntity, nil
}

func (s *Service) UpdateStatus(ctx context.Context, token string, status TokenStatus) *shared_errors.AppError {
	if err := s.tokenRepository.UpdateStatus(ctx, token, status); err != nil {
		return err
	}

	return nil
}
