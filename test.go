package main

import (
	"api-golang/config"
)

func main() {
	// url := config.GetUrlDatabase()
	// fmt.Println(url)

	config.CreateConection()
	config.Ping()

}
