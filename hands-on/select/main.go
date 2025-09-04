package main 

import (
	"fmt"
	"time"
)
func func1(c1 chan<- string) {
	time.Sleep(1 * time.Second)
	c1 <- "hello"
}

func func2(c2 chan<- string) {
	time.Sleep(2 * time.Second)
	c2 <- "world"
}

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func1(c1)
	go func2(c2)

	for range 2 {
		select {
		case msg1 := <-c1 :
			fmt.Println("received :",msg1)
		case msg2 := <-c2 :
			fmt.Println("received :",msg2)
		}
	}
}
