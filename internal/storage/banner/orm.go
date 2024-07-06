package banner

import (
	"github.com/ltphat2204/clothes-store/internal/entity"
	"gorm.io/gorm"
)

type postgreSqlStorage struct {
	storage *gorm.DB
}

func NewPostgresSqlStorage(db *gorm.DB) *postgreSqlStorage {
	db.AutoMigrate(&entity.Banner{})

	return &postgreSqlStorage{
		storage: db,
	}
}
