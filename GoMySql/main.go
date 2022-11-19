package main

import "gomysql/db"

func main(){
	db.Connect()
	db.Close()
}
