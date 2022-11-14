package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	initTime := time.Now()

	channel := make(chan string)

	servers := []string{
		"https://oregoom.com/",
		"https://udemy.com/",
		"https://facebook.com/",
	}

	for _, server := range servers {
		go checkServer(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	timePassed := time.Since(initTime)
	fmt.Println("Time ejecution: ", timePassed)
}

func checkServer(server string, channel chan string) {
	_, err := http.Get(server)
	if err != nil {
		//fmt.Println("Server is not available")
		channel <- server + "Server is not available"
	} else {
		//fmt.Println(server,"Server it´s run")
		channel <- server + "Server it´s run"
	}
}
