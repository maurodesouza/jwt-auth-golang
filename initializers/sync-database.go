package initializers

import "github.com/maurodesouza/jwt-auth-golang/models"

func SyncDatabase() {
	DB.AutoMigrate(models.User{})
}
