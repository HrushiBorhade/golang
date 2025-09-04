package main 

import "fmt"

func main() {
	fmt.Println("Channel Buffering")

	messages := make(chan string, 2)


	messages <- "hello"
	messages <- "world"


	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
