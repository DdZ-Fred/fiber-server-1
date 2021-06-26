package dbUtils

import (
	"github.com/DdZ-Fred/fiber-server-1/models"
	"gorm.io/gorm"
)

func createOrMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
