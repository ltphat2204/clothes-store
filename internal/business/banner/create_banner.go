package banner

import (
	"context"

	"github.com/ltphat2204/clothes-store/internal/dto/banner"
	"github.com/ltphat2204/clothes-store/internal/entity"
)

type createBannerStorage interface {
	CreateBanner(ctx context.Context, banner *banner.BannerRequest) (*entity.Banner, error)
}

type createBannerBusiness struct {
	storage createBannerStorage
}

func NewCreateBannerBusiness(storage createBannerStorage) *createBannerBusiness {
	return &createBannerBusiness{
		storage: storage,
	}
}

func (biz *createBannerBusiness) CreateBanner(ctx context.Context, bannerReq *banner.BannerRequest) (*banner.BannerResult, error) {
	b, err := biz.storage.CreateBanner(ctx, bannerReq)
	if err != nil {
		return nil, err
	}

	var res banner.BannerResult
	res.CopyFrom(b)

	return &res, nil
}