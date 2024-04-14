package service

import (
	"context"
	"github.com/sema0205/avito-backend-assignment-2024/internal/repository"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
	log "github.com/sirupsen/logrus"
)

type AdminService struct {
	repo         repository.Admin
	tokenManager auth.Provider
}

func NewAdminService(repo repository.Admin, tokenManager auth.Provider) Admin {
	return &AdminService{
		repo:         repo,
		tokenManager: tokenManager,
	}
}

func (a *AdminService) SignIn(ctx context.Context, input SignInInput) (string, error) {

	user, err := a.repo.GetByCredentials(ctx, repository.CredentialsInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		log.Errorf("AdminService.SignIn: get by credentials: %v", err)
		return "", err
	}

	accessToken, err := a.tokenManager.NewAdminJWT(user.Id)
	if err != nil {
		log.Errorf("AdminService.SignIn: generate jwt token: %v", err)
		return "", err
	}

	return accessToken, nil
}
