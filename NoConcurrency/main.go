package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){

	initTime := time.Now()
	servers := []string{
		"https://oregoom.com/",
		"https://udemy.com/",
		"https://facebook.com/",
	}

	for _, server := range servers{
		checkServer(server)
	}

	timePassed := time.Since(initTime)
	fmt.Println("Time ejecution: ",timePassed)
}

func checkServer(server string){
	_ , err := http.Get(server)
	if err != nil{
		fmt.Println("Server is not available")
	}else{
		fmt.Println(server,"Server it´s run")
	}
}