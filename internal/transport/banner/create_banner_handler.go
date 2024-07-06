package banner

import (
	"net/http"

	"github.com/gin-gonic/gin"
	bannerDto "github.com/ltphat2204/clothes-store/internal/dto/banner"
	bannerStorage "github.com/ltphat2204/clothes-store/internal/storage/banner"
	bannerBusiness "github.com/ltphat2204/clothes-store/internal/business/banner"
	"github.com/ltphat2204/clothes-store/pkg/common"
	"gorm.io/gorm"
)

func CreateBanner(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var req bannerDto.BannerRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, common.NewSimpleError(err.Error()))
			return
		}

		s := bannerStorage.NewPostgresSqlStorage(db)
		biz := bannerBusiness.NewCreateBannerBusiness(s)

		banner, err := biz.CreateBanner(c.Request.Context(), &req)

		if  err != nil {
			c.JSON(http.StatusInternalServerError, common.NewError(
				"Can not create banner",
				err.Error(),
			))
			return
		}



		c.JSON(http.StatusOK, common.NewSuccessResponse(
			"Banner created successfully",
			banner,
		))
	}
}