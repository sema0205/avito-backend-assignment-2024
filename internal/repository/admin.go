package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
)

type AdminRepo struct {
	db         *pgxpool.Pool
	builder    squirrel.StatementBuilderType
	adminTable string
}

func NewAdminRepo(db *pgxpool.Pool, builder squirrel.StatementBuilderType) Admin {
	return &AdminRepo{
		db:         db,
		builder:    builder,
		adminTable: "service_admin",
	}
}

func (b *AdminRepo) GetByCredentials(ctx context.Context, input CredentialsInput) (domain.Admin, error) {
	selectQuery, args, _ := b.builder.
		Select("*").
		From(b.adminTable).
		Where("username = ? AND password = ?", input.Username, input.Password).
		ToSql()

	rows, err := b.db.Query(ctx, selectQuery, args...)
	defer rows.Close()
	if err != nil {
		return domain.Admin{}, fmt.Errorf("AdminRepo.GetByCredentials - db.Query: %w", err)
	}

	banner, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domain.Admin])
	if err != nil {
		return domain.Admin{}, fmt.Errorf("AdminRepo.GetByCredentials - pgx.CollectOneRow: %w", err)
	}

	return banner, nil
}
