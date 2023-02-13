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

func GetUsers() Users {
	users := Users{}
	// El metodo Find() de GORM recibe un puntero a un slice de usuarios
	// Todos los registros se almacenaran en el slice de usuarios(users)
	db.Find(&users)
	return users
}

func GetUser(id int) *User {
	user := &User{}

	db.Where("id=?", id).First(user)
	return user
}

func (this *User) Save() {
	if this.Id == 0 {
		db.Create(&this)
	} else {
		this.Update()
	}
}

func (this *User) Update() {
	user := User{
		FirstName: this.FirstName,
		LastName:  this.LastName,
		Email:     this.Email,
		Password:  this.Password,
	}

	db.Model(&this).UpdateColumns(user)
}

func (this *User) Delete() {
	db.Delete(&this)
}
