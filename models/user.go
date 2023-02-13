package models

import "errors"

// Lo que va en comillas es el nombre como se identificara en el JSON
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Users []User

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{
		Id:        1,
		FirstName: "Juan",
		LastName:  "Perez",
		Email:     "juan@mail.com",
		Password:  "123456",
	}

	users[user.Id] = user
}

func GetUser(id int) (User, error) {
	if user, ok := users[id]; ok {
		return user, nil
	}
	return User{}, errors.New("User not found")
}

func GetUsers() Users {
	var usersList Users
	for _, user := range users {
		usersList = append(usersList, user)
	}
	return usersList
}

func SaveUser(user User) User {
	user.Id = len(users) + 1
	users[user.Id] = user
	return user
}

func UpdateUser(user User, userResponse User) User {
	user.FirstName = userResponse.FirstName
	user.LastName = userResponse.LastName
	user.Email = userResponse.Email
	user.Password = userResponse.Password
	users[user.Id] = user
	return user
}

func DeleteUser(id int) {
	delete(users, id)
}
