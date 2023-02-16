package models

import (
	"api-golang/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var debug bool

// Este metodo se ejecuta inmediatamente al usar el paquete
// Basicamente cuando el paquete 'models' es importado
func init() {
	CreateConection()
	debug = config.GetDebug()
}

func CreateConection() {

	if GetConnection() != nil {
		return
	}
	url := config.GetUrlDatabase()
	if connection, err := sql.Open("mysql", url); err != nil {
		panic(err.Error())
	} else {
		db = connection
	}
}

func CreateTable(table, schema string) {
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

func existsTable(table string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", table)
	rows, _ := Query(sql)
	return rows.Next()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil && !debug {
		log.Println(err.Error())
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil && !debug {
		log.Println(err.Error())
	}
	return rows, err
}

func InsertData(query string, args ...interface{}) (int64, error) {
	result, err := Exec(query, args...)
	if err != nil {
		return int64(0), err
	} else {
		id, err := result.LastInsertId()
		return id, err
	}
}

func GetConnection() *sql.DB {
	return db
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
