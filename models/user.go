package models

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Lo que va en comillas es el nombre como se identificara en el JSON
type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	createdAt time.Time
}

type Users []User

// Expresion regular para validar el formato de email
var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const UserSchema string = `CREATE TABLE users (
	id INT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
	first_name VARCHAR(50) NULL DEFAULT NULL,
	last_name VARCHAR(50) NULL DEFAULT NULL,
	email VARCHAR(50) NULL DEFAULT NULL UNIQUE,
	password VARCHAR(255) NULL DEFAULT NULL,
	created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (id) USING BTREE)`

// Funcion que crea un nuevo usuario
func NewUser(first_name, last_name, email, password string) (*User, error) {
	user := &User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
	}
	err := user.SetPassword(password)
	if err = user.Valid(); err != nil {
		return &User{}, err
	}
	return user, err
}

func CreateUser(first_name, last_name, email, password string) (*User, error) {
	user, err := NewUser(first_name, last_name, email, password)
	if err != nil {
		return &User{}, err
	}
	err = user.Save()
	return user, err
}

func GetUserByEmail(email string) *User {
	sql := "SELECT id, first_name, last_name, email, password, created_at FROM users WHERE email = ?"
	return GetUser(sql, email)
}

func GetUserById(id int) *User {
	sql := "SELECT id, first_name, last_name, email, password, created_at FROM users WHERE id = ?"
	return GetUser(sql, id)
}

func GetUser(sql string, conditional interface{}) *User {
	user := &User{}
	rows, err := Query(sql, conditional)
	if err != nil {
		return user
	}

	if rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.createdAt)
	}
	return user
}

func GetUsers() Users {

	sql := "SELECT id, first_name, last_name, email, password, created_at FROM users"
	users := Users{}
	rows, _ := Query(sql)

	for rows.Next() {
		user, _ := NewUser("", "", "", "")
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.createdAt)
		users = append(users, *user)
	}
	return users
}

func ValidEmail(email string) error {
	// Devuelve true si el email es valido
	if !emailRegexp.MatchString(email) {
		return errors.New("formato de email invalido")
	}
	return nil
}

func ValidFirstName(first_name string) error {
	if len(first_name) > 50 {
		return errors.New("el nombre es mayor que 50 caracteres")
	}
	return nil
}

func (this *User) Valid() error {
	if err := ValidEmail(this.Email); err != nil {
		return err
	}
	if this.FirstName == "" {
		return errors.New("el nombre es requerido")
	}
	if err := ValidFirstName(this.FirstName); err != nil {
		return err
	}
	if this.LastName == "" {
		return errors.New("el apellido es requerido")
	}

	if this.Password == "" {
		return errors.New("la contraseña es requerida")
	}
	return nil
}

func (this *User) Save() error {

	if this.Id == 0 { // Si el id es 0, es un nuevo usuario. El id se genera automaticamente
		return this.insert()
	} else {
		return this.update()
	}
}

func (this *User) insert() error {
	sql := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)"
	id, err := InsertData(sql, this.FirstName, this.LastName, this.Email, this.Password)
	this.Id = id
	return err
}

func (this *User) update() error {
	sql := "UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ? WHERE id = ?"
	_, err := Exec(sql, this.FirstName, this.LastName, this.Email, this.Password, this.Id)
	return err
}

func (this *User) Delete() error {
	sql := "DELETE FROM users WHERE id = ?"
	_, err := Exec(sql, this.Id)
	return err
}

func (this *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("no es posible cifrar el passwor")
	}
	// hash es un slice de bytes, por eso lo convertimos a string
	this.Password = string(hash)
	return nil
}

func (this *User) GetCreatedAt() time.Time {
	return this.createdAt
}
