package repository

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
	"time"
)

const (
	maxPaginationLimit     = 10
	defaultPaginationLimit = 10
)

type BannerRepo struct {
	db                 *pgxpool.Pool
	builder            squirrel.StatementBuilderType
	bannerTable        string
	bannerFeatureTable string
	bannerTagTable     string
}

func NewBannerRepo(db *pgxpool.Pool, builder squirrel.StatementBuilderType) Banner {
	return &BannerRepo{
		db:                 db,
		builder:            builder,
		bannerTable:        "banner",
		bannerFeatureTable: "banner_feature_item",
		bannerTagTable:     "banner_tag_item",
	}
}

func (b *BannerRepo) Create(ctx context.Context, banner domain.Banner) (int, error) {
	tx, err := b.db.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("BannerRepo.Create - db.Begin: %w", err)
	}
	defer tx.Rollback(ctx)

	insertQuery, args, _ := b.builder.
		Insert(b.bannerTable).
		Columns("content", "is_active").
		Values(string(banner.Content), banner.IsActive).
		Suffix("RETURNING id").
		ToSql()

	var bannerId int
	err = tx.QueryRow(ctx, insertQuery, args...).Scan(&bannerId)
	if err != nil {
		return 0, fmt.Errorf("BannerRepo.insertBanner - tx.QueryRow: %w", err)
	}

	insertQuery, args, _ = b.builder.
		Insert(b.bannerFeatureTable).
		Columns("banner_id", "feature_id").
		Values(bannerId, banner.FeatureId).
		ToSql()

	_, err = tx.Exec(ctx, insertQuery, args...)
	if err != nil {
		return 0, fmt.Errorf("BannerRepo.insertFeatureBanner - tx.Exec: %w", err)
	}

	queryBuilder := b.builder.
		Insert(b.bannerTagTable).
		Columns("banner_id", "tag_id")

	for _, tagId := range banner.TagsIds {
		queryBuilder = queryBuilder.Values(bannerId, tagId)
	}

	insertQuery, args, _ = queryBuilder.ToSql()
	_, err = tx.Exec(ctx, insertQuery, args...)
	if err != nil {
		return 0, fmt.Errorf("BannerRepo.insertTagsBanner - tx.Exec: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("BannerRepo.Create - tx.Commit: %w", err)
	}

	return bannerId, nil
}

func (b *BannerRepo) Update(ctx context.Context, input UpdateBannerInput) error {

	tx, err := b.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("BannerRepo.Update - db.Begin: %w", err)
	}
	defer tx.Rollback(ctx)

	if input.Content != nil {
		updateQuery, args, _ := b.builder.
			Update(b.bannerTable).
			Set("content", *input.Content).
			Set("updated_at", time.Now()).
			Where("id = ?", input.Id).
			ToSql()

		_, err = tx.Exec(ctx, updateQuery, args...)
		if err != nil {
			return fmt.Errorf("BannerRepo.UpdateContent - tx.Exec: %w", err)
		}
	}

	if input.FeatureId != nil {
		updateQuery, args, _ := b.builder.
			Update(b.bannerFeatureTable).
			Set("feature_id", *input.FeatureId).
			Where("banner_id = ?", input.Id).
			ToSql()

		_, err = tx.Exec(ctx, updateQuery, args...)
		if err != nil {
			return fmt.Errorf("BannerRepo.UpdateFeature - tx.Exec: %w", err)
		}

		updateQuery, args, _ = b.builder.
			Update(b.bannerTable).
			Set("updated_at", time.Now()).
			Where("id = ?", input.Id).
			ToSql()

		_, err = tx.Exec(ctx, updateQuery, args...)
		if err != nil {
			return fmt.Errorf("BannerRepo.UpdateBanner - tx.Exec: %w", err)
		}
	}

	if len(input.TagIds) != 0 {
		deleteQuery, args, _ := b.builder.
			Delete(b.bannerTagTable).
			Where("banner_id = ?", input.Id).
			ToSql()

		_, err = tx.Exec(ctx, deleteQuery, args...)
		if err != nil {
			return fmt.Errorf("BannerRepo.DeleteOldTags - tx.Exec: %w", err)
		}

		tagQueryBuilder := b.builder.
			Insert(b.bannerTagTable).
			Columns("banner_id", "tag_id")

		for _, tagId := range input.TagIds {
			if tagId != nil {
				tagQueryBuilder = tagQueryBuilder.Values(input.Id, *tagId)
			}
		}

		insertQuery, args, _ := tagQueryBuilder.ToSql()
		_, err = tx.Exec(ctx, insertQuery, args...)
		if err != nil {
			return fmt.Errorf("BannerRepo.InsertNewTags - tx.Exec: %w", err)
		}

		updateQuery, args, _ := b.builder.
			Update(b.bannerTable).
			Set("updated_at", time.Now()).
			Where("id = ?", input.Id).
			ToSql()

		_, err = tx.Exec(ctx, updateQuery, args...)
		if err != nil {
			return fmt.Errorf("BannerRepo.UpdateBanner - tx.Exec: %w", err)
		}
	}

	if input.IsActive != nil {
		updateQuery, args, _ := b.builder.
			Update(b.bannerTable).
			Set("is_active", *input.IsActive).
			Set("updated_at", time.Now()).
			Where("id = ?", input.Id).
			ToSql()

		_, err = tx.Exec(ctx, updateQuery, args...)
		if err != nil {
			return fmt.Errorf("BannerRepo.UpdateIsActive - tx.Exec: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("BannerRepo.Update - tx.Commit: %w", err)
	}

	return nil
}

func (b *BannerRepo) Delete(ctx context.Context, id int) error {
	insertQuery, args, _ := b.builder.
		Delete(b.bannerTable).
		Where("id = ?", id).
		ToSql()

	_, err := b.db.Exec(ctx, insertQuery, args...)
	if err != nil {
		return fmt.Errorf("BannerRepo.Delete - db.Exec: %w", err)
	}

	return nil
}

func (b *BannerRepo) GetByTagIdAndFeatureId(ctx context.Context, input GetUserBannerInput) (domain.Banner, error) {
	selectQuery, args, _ := b.builder.
		Select("b.content").
		From(fmt.Sprintf("%s AS b", b.bannerTable)).
		InnerJoin(fmt.Sprintf("%s AS bf ON b.id = bf.feature_id", b.bannerFeatureTable)).
		InnerJoin(fmt.Sprintf("%s AS bt ON b.id = bt.tag_id", b.bannerTagTable)).
		Where("feature_id = ? AND tag_id = ? AND is_active = ?", input.FeatureId, input.TagId, true).
		ToSql()

	rows, err := b.db.Query(ctx, selectQuery, args...)
	defer rows.Close()
	if err != nil {
		return domain.Banner{}, fmt.Errorf("BannerRepo.GetByTagIdAndFeatureId - db.Query: %w", err)
	}

	banner, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domain.Banner])
	if err != nil {
		return domain.Banner{}, fmt.Errorf("BannerRepo.GetByTagIdAndFeatureId - pgx.CollectOneRow: %w", err)
	}

	return banner, nil
}

func (b *BannerRepo) GetFilteredBanners(ctx context.Context, input GetFilteredBannerInput) ([]domain.Banner, error) {
	if input.Limit > maxPaginationLimit {
		input.Limit = maxPaginationLimit
	}
	if input.Limit == 0 {
		input.Limit = defaultPaginationLimit
	}

	selectQuery := b.builder.
		Select("b.id", "b.content", "b.is_active", "b.created_at", "b.updated_at", "array_agg(bt.tag_id) AS tag_ids", "bf.feature_id").
		From(fmt.Sprintf("%s AS b", b.bannerTable)).
		LeftJoin(fmt.Sprintf("%s AS bt ON bt.banner_id = b.id", b.bannerTagTable)).
		LeftJoin(fmt.Sprintf("%s AS bf ON bf.banner_id = b.id", b.bannerFeatureTable))

	if input.FeatureId != nil {
		selectQuery = selectQuery.
			Where(squirrel.Eq{"bf.feature_id": *input.FeatureId})
	}

	if input.TagID != nil {
		selectQuery = selectQuery.
			Where(squirrel.Eq{"bt.tag_id": *input.TagID})
	}

	sqlStr, args, _ := selectQuery.
		GroupBy("b.id", "bf.feature_id").
		OrderBy("b.created_at DESC").
		Limit(uint64(input.Limit)).
		Offset(uint64(input.Offset)).
		ToSql()

	rows, err := b.db.Query(ctx, sqlStr, args...)
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("BannerRepo.GetFilteredBanners - db.Query: %w", err)
	}

	filteredBanners, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[domain.Banner])
	if err != nil {
		return nil, fmt.Errorf("BannerRepo.GetFilteredBanners - pgx.CollectOneRow: %w", err)
	}

	return filteredBanners, nil
}
