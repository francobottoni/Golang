package models

import "gomysql/db"

type User struct {
	Id       int64
	UserName string
	Password string
	Email    string
}

const UserSchema string = `CREATE TABLE users (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(30) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(50),
    create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP )`

type Users []User

// Construir usuario
func NewUser(username, password, email string) *User {
	user := &User{UserName: username, Password: password, Email: email}
	return user
}

// Crear usuario e inserta
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// Inserta Registro
func (user *User) insert() {
	sql := "INSERT users SET username =?, password=?, email=?"
	result, _ := db.Exec(sql, user.UserName, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// Listar todos los registros
func ListUsers() Users {
	sql := "SELECT id,username, password,email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

// Obtener un solo registro
func GetUser(id int) *User {
	sql := "SELECT id,username, password,email FROM users where id=?"
	rows, _ := db.Query(sql, id)
	user := &User{}
	for rows.Next() {
		rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
	}

	return user
}

// Actualizar registro
func (user *User) update() {
	sql := "UPDATE users SET username =?, password=?, email=? where id=?"
	db.Exec(sql, user.UserName, user.Password, user.Email, user.Id)
}

// Guarda o edita registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// Eliminar un registro
func (user *User) Delete() {
	sql := "DELETE FROM users where id=?"
	db.Exec(sql, user.Id)
}
