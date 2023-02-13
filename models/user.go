package models

import "api-golang/config"

// Lo que va en comillas es el nombre como se identificara en el JSON
type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Users []User

const UserSchema string = `CREATE TABLE 'users' (
			'id' INT(11) NOT NULL AUTO_INCREMENT,
			'first_name' VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
			'last_name' VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
			'email' VARCHAR(50) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
			'password' VARCHAR(255) NULL DEFAULT NULL COLLATE 'latin1_swedish_ci',
			'created_at' TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY ('id') USING BTREE,
			UNIQUE INDEX 'email' ('email') USING BTREE
		)
		COLLATE='latin1_swedish_ci'
		ENGINE=InnoDB`

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

func GetUser(id int) *User {
	user := NewUser("", "", "", "")
	sql := "SELECT id, first_name, last_name, email, password FROM users WHERE id = ?"
	rows, _ := config.Query(sql, id)

	if rows.Next() {
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	}
	return user
}

func GetUsers() Users {

	sql := "SELECT id, first_name, last_name, email, password FROM users"
	users := Users{}
	rows, _ := config.Query(sql)

	for rows.Next() {
		user := NewUser("", "", "", "")
		rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		users = append(users, *user)
	}
	return users
}

func (this *User) Save() {

	if this.Id == 0 { // Si el id es 0, es un nuevo usuario. El id se genera automaticamente
		this.insert()
	} else {
		this.update()
	}
}

func (this *User) insert() {
	sql := "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)"
	result, _ := config.Exec(sql, this.FirstName, this.LastName, this.Email, this.Password)
	this.Id, _ = result.LastInsertId() // int64
}

func (this *User) update() {
	sql := "UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ? WHERE id = ?"
	config.Exec(sql, this.FirstName, this.LastName, this.Email, this.Password, this.Id)
}

func (this *User) Delete() {
	sql := "DELETE FROM users WHERE id = ?"
	config.Exec(sql, this.Id)
}
