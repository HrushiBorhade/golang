
package main 

import "fmt"

func pings(ping chan<- string, msg string) {
	ping <- msg
}
func pongs(ping <-chan string, pong chan<-string){
	msg := <- ping
	pong <- msg
}
func main() {
	ping := make(chan string, 1)
	pong := make(chan string, 1)

	pings(ping,"passed message")
	pongs(ping, pong)
	fmt.Println(<-pong)
}
