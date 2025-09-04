package main

import (
	"fmt"
	"time"
)


func main() {

	msg := make(chan string)


	go func(message string) {
		msg <- message
		time.Sleep(time.Second)
		msg <- "this wont get printed as process is closed"
	}("ping")

	messages := <- msg
	fmt.Println(messages)
}
