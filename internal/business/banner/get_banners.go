package banner

import (
	"context"

	"github.com/ltphat2204/clothes-store/internal/dto/banner"
	"github.com/ltphat2204/clothes-store/internal/entity"
)

type getBannersStorage interface {
	GetBanners(ctx context.Context) ([]entity.Banner, error)
}

type getBannersBusiness struct {
	storage getBannersStorage
}

func NewGetBannersBusiness(storage getBannersStorage) *getBannersBusiness {
	return &getBannersBusiness{
		storage: storage,
	}
}

func (biz *getBannersBusiness) GetBanners(ctx context.Context) ([]banner.BannerResult, error) {
	banners, err := biz.storage.GetBanners(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]banner.BannerResult, 0)
	for _, b := range banners {
		var res banner.BannerResult
		res.CopyFrom(&b)
		results = append(results, res)
	}

	return results, nil
}