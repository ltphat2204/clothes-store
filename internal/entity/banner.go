package entity

import "gorm.io/gorm"

type Banner struct {
	gorm.Model
	Image       string `gorm:"type:varchar(255)"`
	ImageSmall  string `gorm:"type:varchar(255)"`
	RedirectUrl string `gorm:"type:varchar(255)"`
}

func (Banner) TableName() string {
	return "banners"
}
