package test

import (
	"api-golang/migrations"
	"api-golang/models"
	"fmt"
	"os"
	"testing"
)

// Esta funcion nos permite ejecutar todas las pruebas unitarias de una sola vez
// Todas las pruebas qeu se encuentren dentro del paquete test, osea que esten en la misma carpeta. No solamente del mismo archivo
// Esta funcion TestMain SOLO debe estar una SOLA vez en toda la carpeta que barque el paquete
func TestMain(m *testing.M) {
	beforeTest()
	result := m.Run()
	afterTest()
	os.Exit(result)
}

func beforeTest() {
	fmt.Println(">>> Iniciando pruebas unitarias")
	models.CreateConection()
	migrations.Migrate()
	createDefaultUser()
}

func createDefaultUser() {
	sql := fmt.Sprintf("INSERT users SET id='%d', first_name='%s', last_name='%s', email='%s', password='%s', created_at='%s';", id, firstName, lastName, email, passwordHashed, createdAt)
	_, err := models.Exec(sql)
	if err != nil {
		panic(err.Error())
	}

	// Esta variable fue declarada en el archivo user_test.go, aqui la estamos asignando los valores.
	user = &models.User{Id: id, FirstName: firstName, LastName: lastName, Email: email, Password: passwordHashed}
}

func afterTest() {
	fmt.Println(">>> Finalizando pruebas unitarias")
	models.CloseConection()
}
