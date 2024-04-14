package service

import (
	"context"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
	"github.com/sema0205/avito-backend-assignment-2024/internal/repository"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/cache"
)

type UpdateBannerInput struct {
	Id        int
	Content   *string
	TagIds    []*int
	FeatureId *int
	IsActive  *bool
}

type GetUserBannerInput struct {
	TagId           int
	FeatureId       int
	UseLastRevision bool
}

type GetFilteredBannerInput struct {
	FeatureId *int
	TagID     *int
	Limit     int
	Offset    int
}

type Banner interface {
	Create(ctx context.Context, banner domain.Banner) (int, error)
	Update(ctx context.Context, input UpdateBannerInput) error
	Delete(ctx context.Context, id int) error
	GetByTagIdAndFeatureId(ctx context.Context, input GetUserBannerInput) (domain.Banner, error)
	GetFilteredBanners(ctx context.Context, input GetFilteredBannerInput) ([]domain.Banner, error)
}

type SignInInput struct {
	Username string
	Password string
}

type Admin interface {
	SignIn(ctx context.Context, input SignInInput) (string, error)
}

type User interface {
	SignIn(ctx context.Context, input SignInInput) (string, error)
}

type Deps struct {
	Repos         *repository.Repositories
	TokenManager  auth.Provider
	CacheProvider cache.Provider
}

type Services struct {
	Admin  Admin
	User   User
	Banner Banner
}

func NewServices(deps Deps) *Services {
	return &Services{
		Admin:  NewAdminService(deps.Repos.Admin, deps.TokenManager),
		User:   NewUserService(deps.Repos.User, deps.TokenManager),
		Banner: NewBannerService(deps.Repos.Banner, deps.CacheProvider),
	}
}
