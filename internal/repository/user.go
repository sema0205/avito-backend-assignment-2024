package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
)

type UserRepo struct {
	db        *pgxpool.Pool
	builder   squirrel.StatementBuilderType
	userTable string
}

func NewUserRepo(db *pgxpool.Pool, builder squirrel.StatementBuilderType) User {
	return &UserRepo{
		db:        db,
		builder:   builder,
		userTable: "service_user",
	}
}

func (b *UserRepo) GetByCredentials(ctx context.Context, input CredentialsInput) (domain.User, error) {
	selectQuery, args, _ := b.builder.
		Select("*").
		From(b.userTable).
		Where("username = ? AND password = ?", input.Username, input.Password).
		ToSql()

	rows, err := b.db.Query(ctx, selectQuery, args...)
	defer rows.Close()
	if err != nil {
		return domain.User{}, fmt.Errorf("UserRepo.GetByCredentials - db.Query: %w", err)
	}

	banner, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domain.User])
	if err != nil {
		return domain.User{}, fmt.Errorf("UserRepo.GetByCredentials - pgx.CollectOneRow: %w", err)
	}

	return banner, nil
}
