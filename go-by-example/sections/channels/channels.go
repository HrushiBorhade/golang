package channels

import "fmt"

func Execute() {
	messages := make(chan string)

	go func() {
		messages <- "Ping"
	}()

	msg := <-messages

	fmt.Println(msg)
}
