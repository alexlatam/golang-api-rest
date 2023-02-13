package migrations

import (
	"api-golang/models"
)

func Migrate() {
	models.CreateConection()
	createTables()
	models.CloseConection()
}

func createTables() {
	models.CreateTable("users", models.UserSchema)
	// createTable("posts", models.userSchema)
	// createTable("categories", models.userSchema)
}
