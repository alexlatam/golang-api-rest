package main

import (
	"api-golang/orm"
	"fmt"
)

func main() {
	orm.CreateConnection()
	orm.CreateTables()

	user := orm.NewUser("Juan", "Perez", "test@mail.com", "123456")
	user.Save()

	fmt.Println(user)

}
