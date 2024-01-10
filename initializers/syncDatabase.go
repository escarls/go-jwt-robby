package initializers

import (
	"github.com/escarls/go-jwt-robby/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
