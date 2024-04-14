package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
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

type CredentialsInput struct {
	Username string
	Password string
}

type Admin interface {
	GetByCredentials(ctx context.Context, input CredentialsInput) (domain.Admin, error)
}

type User interface {
	GetByCredentials(ctx context.Context, input CredentialsInput) (domain.User, error)
}

type Repositories struct {
	Banner Banner
	User   User
	Admin  Admin
}

func NewRepositories(db *pgxpool.Pool, qb squirrel.StatementBuilderType) *Repositories {
	return &Repositories{
		Banner: NewBannerRepo(db, qb),
		User:   NewUserRepo(db, qb),
		Admin:  NewAdminRepo(db, qb),
	}
}
