package main

import (
	"api-golang/orm"
	"fmt"
)

func main() {
	orm.CreateConnection()
	orm.CreateTables()

	fmt.Println("------------")
	user := orm.NewUser("Juan", "Perez", "test@mail.com", "123456")
	user.Save()
	fmt.Println(user)
	fmt.Println("------------")

	users := orm.GetUsers()
	fmt.Println(users)

	fmt.Println("------------")

	userGet := orm.GetUser(1)
	fmt.Println(userGet)
	fmt.Println("------------")

	userGet.FirstName = "Juan Carlos"
	userGet.LastName = "Snchez"
	userGet.Email = "sanchaz@mail.com"
	userGet.Save()

	userGet.Delete()

}
