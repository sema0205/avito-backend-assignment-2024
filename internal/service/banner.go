package service

import (
	"context"
	"fmt"
	"github.com/sema0205/avito-backend-assignment-2024/internal/domain"
	"github.com/sema0205/avito-backend-assignment-2024/internal/repository"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/cache"
	log "github.com/sirupsen/logrus"
)

type BannerService struct {
	repo  repository.Banner
	cache cache.Provider
}

func NewBannerService(repo repository.Banner, cache cache.Provider) Banner {
	return &BannerService{
		repo:  repo,
		cache: cache,
	}
}

func (b *BannerService) Create(ctx context.Context, banner domain.Banner) (int, error) {
	return b.repo.Create(ctx, banner)
}

func (b *BannerService) Update(ctx context.Context, input UpdateBannerInput) error {
	return b.repo.Update(ctx, repository.UpdateBannerInput{
		Id:        input.Id,
		Content:   input.Content,
		TagIds:    input.TagIds,
		FeatureId: input.FeatureId,
		IsActive:  input.IsActive,
	})
}

func (b *BannerService) Delete(ctx context.Context, id int) error {
	return b.repo.Delete(ctx, id)
}

func (b *BannerService) GetByTagIdAndFeatureId(ctx context.Context, input GetUserBannerInput) (domain.Banner, error) {
	cacheKey := fmt.Sprintf("banner-%d-%d", input.TagId, input.FeatureId)
	if input.UseLastRevision {

		if cachedBanner, found := b.cache.Get(cacheKey); found {
			return cachedBanner.(domain.Banner), nil
		}
	}

	dbBanner, err := b.repo.GetByTagIdAndFeatureId(ctx, repository.GetUserBannerInput{
		TagId:           input.TagId,
		FeatureId:       input.FeatureId,
		UseLastRevision: input.UseLastRevision,
	})
	if err != nil {
		log.Errorf("BannerService.GetByTagIdAndFeatureId: repository call: %v", err)
		return domain.Banner{}, err
	}

	b.cache.Set(cacheKey, dbBanner)

	return dbBanner, nil
}

func (b *BannerService) GetFilteredBanners(ctx context.Context, input GetFilteredBannerInput) ([]domain.Banner, error) {
	return b.repo.GetFilteredBanners(ctx, repository.GetFilteredBannerInput{
		FeatureId: input.FeatureId,
		TagID:     input.TagID,
		Limit:     input.Limit,
		Offset:    input.Offset,
	})
}
