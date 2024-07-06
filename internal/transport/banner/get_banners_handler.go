package banner

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bannerStorage "github.com/ltphat2204/clothes-store/internal/storage/banner"
	bannerBusiness "github.com/ltphat2204/clothes-store/internal/business/banner"
	"github.com/ltphat2204/clothes-store/pkg/common"
	"gorm.io/gorm"
)

func GetBanners(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		s := bannerStorage.NewPostgresSqlStorage(db)
		biz := bannerBusiness.NewGetBannersBusiness(s)

		banners, err := biz.GetBanners(c.Request.Context())

		if  err != nil {
			c.JSON(http.StatusInternalServerError, common.NewError(
				"Can not get banners",
				err.Error(),
			))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(
			"Banners retrieved successfully",
			banners,
		))
	}
}