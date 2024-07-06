package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ltphat2204/clothes-store/internal/transport/banner"
	"github.com/ltphat2204/clothes-store/pkg/database"
)

func createBannerRoutes(r *gin.RouterGroup) {
	bannerRoutes := r.Group("/banner")
	{
		bannerRoutes.POST("", banner.CreateBanner(database.Database))
	}
}