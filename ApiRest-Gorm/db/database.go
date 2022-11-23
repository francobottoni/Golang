package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// username:password@tcp(localhost:3306)/database
const url = "root:9338QOBU@tcp(localhost:3306)/goweb_db"

var db *sql.DB

// Creamos la conxion
func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = connection
}

// Cerramos la conexion
func Close() {
	db.Close()
}

// Verificamos la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Verifica si una tabla existe o no
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	return rows.Next()
}

// Crea una tabla
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := Exec(schema)
		if err != nil {
			panic(err)
		}
	}
}

// Reiniciar el registro de una tabla
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// Polimorfismo de Exec
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()

	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
