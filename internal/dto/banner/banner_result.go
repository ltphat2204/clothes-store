package banner

import "github.com/ltphat2204/clothes-store/internal/entity"

type BannerResult struct {
	Image       string `json:"image"`
	ImageSmall  string `json:"image_small"`
	RedirectUrl string `json:"redirect_url"`
}

func (banner *BannerResult) CopyFrom(other *entity.Banner) { 
	banner.Image = other.Image
	banner.ImageSmall = other.ImageSmall
	banner.RedirectUrl = other.RedirectUrl
}