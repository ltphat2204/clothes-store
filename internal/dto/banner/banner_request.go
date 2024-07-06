package banner

import "github.com/ltphat2204/clothes-store/internal/entity"

type BannerRequest struct {
	Image       string `binding:"required" json:"image"`
	ImageSmall  string `binding:"required" json:"image_small"`
	RedirectUrl string `binding:"required" json:"redirect_url"`
}

func (banner *BannerRequest) CopyTo(other *entity.Banner) { 
	other.Image = banner.Image
	other.ImageSmall = banner.ImageSmall
	other.RedirectUrl = banner.RedirectUrl
}