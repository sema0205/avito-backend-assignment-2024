package service

import (
	"context"
	"github.com/sema0205/avito-backend-assignment-2024/internal/repository"
	log "github.com/sirupsen/logrus"

	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
)

type UserService struct {
	repo         repository.User
	tokenManager auth.Provider
}

func NewUserService(repo repository.User, tokenManager auth.Provider) User {
	return &UserService{
		repo:         repo,
		tokenManager: tokenManager,
	}
}

func (u *UserService) SignIn(ctx context.Context, input SignInInput) (string, error) {

	user, err := u.repo.GetByCredentials(ctx, repository.CredentialsInput{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		log.Errorf("UserService.SignIn: get by credentials: %v", err)
		return "", err
	}

	accessToken, err := u.tokenManager.NewUserJWT(user.Id)
	if err != nil {
		log.Errorf("UserService.SignIn: generate jwt token: %v", err)
		return "", err
	}

	return accessToken, nil
}
