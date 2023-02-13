package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const username string = "root"
const password string = ""
const host string = "localhost"
const port int = 3306
const database string = "api-golang"

func CreateTables() {
	createTable("users", models.userSchema)
	// createTable("posts", models.userSchema)
	// createTable("categories", models.userSchema)
}

func createTable(table, schema string) {
	if !existsTable(table) {
		Exec(schema)
	} else {
		truncateTable(table)
	}
}

func truncateTable(table string) {
	sql := fmt.Sprintf("TRUNCATE TABLE %s", table)
	Exec(sql)
}

func CreateConection() {
	connection, err := sql.Open("mysql", generateURL())
	if err != nil {
		panic(err.Error())
	} else {
		db = connection
	}
}

func existsTable(table string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", table)
	rows, _ := Query(sql)
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		log.Println(err.Error())
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println(err.Error())
	}
	return rows, err
}

// Funcion que verifica que la conexion a la base de datos este viva(abierta)
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
}

func CloseConection() {
	db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
}
