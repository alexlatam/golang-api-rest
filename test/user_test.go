package test

import (
	"api-golang/models"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var user *models.User

const (
	id             = 1
	firstName      = "alexis"
	lastName       = "Montilla"
	email          = "alexis@test.com"
	password       = "1234"
	passwordHashed = "$2a$10$AgGD6ypPhikajC5yE8nAbuuyVVOpWz2pFWalDkUfL.GvCUE4TtbtO"
	createdAt      = "1991-10-26 00:00:00"
)

func TestNewUser(t *testing.T) {
	_, err := models.NewUser(firstName, lastName, email, password)
	if err != nil {
		t.Error("El usuario no se creo correctamente", err)
	}
	// if !equalsUser(user) {
	// 	t.Error("El usuario no se creo correctamente")
	// }
}

func TestPassword(t *testing.T) {
	user, _ := models.NewUser(firstName, lastName, email, password)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		t.Error("El password no coincide")
	}
}

func TestValidEmail(t *testing.T) {
	if err := models.ValidEmail(email); err != nil {
		t.Error("El formato del email no es valido")
	}
}

func TestInvalidEmail(t *testing.T) {
	if err := models.ValidEmail("invalidFormat"); err == nil {
		t.Error("El formato del email no es valido")
	}
}

func TestFirstNameLength(t *testing.T) {
	newFirstName := firstName
	for i := 0; i < 10; i++ {
		newFirstName += newFirstName
	}
	_, err := models.NewUser(newFirstName, lastName, email, password)
	if err != nil {
		t.Error("es posible crear un usuario en BD con un nombre mayor a 50 caracteres")
	}
}

func TestSave(t *testing.T) {
	user, _ := models.NewUser(firstName, lastName, randomEmail(), password)
	if err := user.Save(); err != nil {
		t.Error("El usuario no se guardo correctamente", err)
	}
}

func TestCreateUser(t *testing.T) {
	_, err := models.CreateUser(firstName, lastName, randomEmail(), password)
	if err != nil {
		t.Error("El usuario no se creo correctamente", err)
	}
}

func TestUniqueEmail(t *testing.T) {
	_, err := models.CreateUser(firstName, lastName, email, password)
	if err == nil { // Si el error es nil, es porque no se genero ningun error, es decir, que se creo el usuario correctamente
		t.Error("Alerta! Es posible crear dos usuarios con el mismo email!")
	}
}

// Este test valida que el error que genero la BD al intentar crear un usuario con un email duplicado sea el correcto
// Basicamente la BD puede generar errores por muchos motivos al momento de insertar un nuevo registro
// Aqui nos aseguramos que el error que se esta generando es realmente por data duplicada(en esta caso email)
func TestDuplicateEmail(t *testing.T) {
	_, err := models.CreateUser(firstName, lastName, email, password)
	mysqlErrorMessage := fmt.Sprintf("Error 1062: Duplicate entry '%s' for key 'email'", email)
	if err.Error() != mysqlErrorMessage {
		t.Error("Alerta! Es posible tener un email duplicado en BD!")
	}
}

func TestGetUser(t *testing.T) {
	user := models.GetUserById(id)
	if !equalsUser(user) || !equalsCreatedAt(user.GetCreatedAt()) {
		t.Error("No es posible obtener el usuario")
	}
}

func TestGetUsers(t *testing.T) {
	users := models.GetUsers()
	if len(users) == 0 {
		t.Error("No es posible obtener los usuarios")
	}
}

func TestDeleteUser(t *testing.T) {
	if err := user.Delete(); err != nil {
		t.Error("No es posible eliminar el usuario")
	}
}

func equalsCreatedAt(date time.Time) bool {
	format := "2020-01-01 00:00:00"
	// Parseamos el string createdAt a un tipo de dato time.Time
	dateFormated, _ := time.Parse(format, createdAt)
	return date == dateFormated
}

func equalsUser(user *models.User) bool {
	return user.FirstName == firstName && user.LastName == lastName && user.Email == email
}

func randomEmail() string {
	return fmt.Sprintf("%s_%d@mail.com", firstName, rand.Intn(1000))
}
