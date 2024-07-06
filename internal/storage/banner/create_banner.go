package banner

import (
	"context"
	"github.com/ltphat2204/clothes-store/internal/entity"
	"github.com/ltphat2204/clothes-store/internal/dto/banner"
)

func (s *postgreSqlStorage) CreateBanner(ctx context.Context, banner *banner.BannerRequest) (*entity.Banner, error) {
	bannerEntity := entity.Banner{}
	banner.CopyTo(&bannerEntity)

	if err := s.storage.Create(&bannerEntity).Error; err != nil {
		return nil, err
	}

	return &bannerEntity, nil
}