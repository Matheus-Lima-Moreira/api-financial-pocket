package token

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/iam/identity/user"
	emails "github.com/Matheus-Lima-Moreira/financial-pocket/internal/notifications/emails"
	shared_errors "github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/errors"
)

type Service struct {
	tokenRepository Repository
	emailSender     emails.EmailSender
	userRepository  user.Repository
	verifyBaseURL   string
}

func NewService(tokenRepository Repository, emailSender emails.EmailSender, userRepository user.Repository, verifyBaseURL string) *Service {
	return &Service{
		tokenRepository: tokenRepository,
		emailSender:     emailSender,
		userRepository:  userRepository,
		verifyBaseURL:   verifyBaseURL,
	}
}

func (s *Service) SendVerificationEmail(ctx context.Context, email string) *shared_errors.AppError {

	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	token, err := GenerateToken()
	if err != nil {
		return shared_errors.NewBadRequest(err.Error())
	}

	tokenEntity := NewTokenEntity(TokenResourceVerifyEmail, token, strconv.Itoa(int(user.ID)), map[string]any{
		"email": email,
		"name":  user.Name,
	})

	if err := s.tokenRepository.Create(ctx, tokenEntity); err != nil {
		return err
	}

	verifyURL := s.verifyBaseURL + "?token=" + url.QueryEscape(tokenEntity.Token)

	if err := s.emailSender.SendVerifyEmail(ctx, email, user.Name, verifyURL); err != nil {
		return shared_errors.NewBadRequest("token.verify_email_send_failed")
	}

	return nil
}

func (s *Service) ResetPassword(ctx context.Context, email string) *shared_errors.AppError {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	token, err := GenerateToken()
	if err != nil {
		return shared_errors.NewBadRequest(err.Error())
	}

	tokenEntity := NewTokenEntity(TokenResourceResetPassword, token, strconv.Itoa(int(user.ID)), map[string]any{
		"email": email,
		"name":  user.Name,
	})

	if err := s.tokenRepository.Create(ctx, tokenEntity); err != nil {
		return err
	}

	resetPasswordURL := s.verifyBaseURL + "?token=" + url.QueryEscape(tokenEntity.Token)
	if err := s.emailSender.SendResetPasswordEmail(ctx, email, user.Name, resetPasswordURL); err != nil {
		return shared_errors.NewBadRequest("token.reset_password_send_failed")
	}

	return nil
}

func (s *Service) VerifyEmail(ctx context.Context, verifyToken string) *shared_errors.AppError {
	tokenEntity, err := s.tokenRepository.FindByToken(ctx, verifyToken)
	if err != nil {
		return err
	}

	if tokenEntity.Resource != TokenResourceVerifyEmail {
		return shared_errors.NewUnauthorized("error.invalid_token")
	}

	if tokenEntity.Status != TokenStatusActive {
		return shared_errors.NewUnauthorized("error.invalid_token")
	}

	if tokenEntity.ExpiresAt.Before(time.Now()) {
		return shared_errors.NewUnauthorized("error.expired_token")
	}

	if err := s.tokenRepository.UpdateStatus(ctx, verifyToken, TokenStatusUsed); err != nil {
		return err
	}

	parsedID, parseErr := strconv.Atoi(tokenEntity.ReferenceID)
	if parseErr != nil {
		return shared_errors.NewUnauthorized("error.invalid_token")
	}

	if err := s.userRepository.SetEmailVerified(ctx, uint(parsedID), true); err != nil {
		return err
	}

	return nil
}
