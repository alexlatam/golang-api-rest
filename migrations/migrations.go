package migrations

import (
	"api-golang/models"
)

func Migrate() {
	models.CreateConection()
	createTables()
}

func createTables() {
	models.CreateTable("users", models.UserSchema)
	// createTable("posts", models.userSchema)
	// createTable("categories", models.userSchema)
}
