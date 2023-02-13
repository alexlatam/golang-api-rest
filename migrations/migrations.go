package migrations

import (
	"api-golang/config"
	"api-golang/models"
)

func Migrate() {
	config.CreateConection()
	createTables()
	config.CloseConection()
}

func createTables() {
	config.CreateTable("users", models.UserSchema)
	// createTable("posts", models.userSchema)
	// createTable("categories", models.userSchema)
}
