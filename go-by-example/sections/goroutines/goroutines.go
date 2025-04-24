package goroutines

import (
	"fmt"
	"time"
)

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, " : ", i)
	}
}

func Execute() {
	f("Main thread")

	go f("first go routine")

	go func(str string) {
		fmt.Println(str)
	}("Second go routine")

	time.Sleep(time.Second)
	fmt.Println("Done here")
}
