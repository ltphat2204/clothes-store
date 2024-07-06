package banner

import (
	"context"
	"github.com/ltphat2204/clothes-store/internal/entity"
)

func (s *postgreSqlStorage) GetBanners(ctx context.Context) ([]entity.Banner, error) {
	var banners []entity.Banner
	if err := s.storage.Find(&banners).Order("created_at desc").Error; err != nil {
		return nil, err
	}

	return banners, nil
}