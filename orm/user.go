package orm

import "time"

type User struct {
	Id        int64     `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Users []User

//  Funcion que crea un nuevo usuario
func NewUser(first_name, last_name, email, password string) *User {
	user := &User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Password:  password,
	}
	return user
}

func CreateUser(first_name, last_name, email, password string) *User {
	user := NewUser(first_name, last_name, email, password)
	user.Save()
	return user
}

func (this *User) Save() {
	db.Create(&this)
}
