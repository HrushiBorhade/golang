package main

import "fmt"

func main() {

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Printf("I'm a bool with value: %v\n", t)
		case int:
			fmt.Printf("I'm an int with value: %d\n", t)
		default:
			fmt.Printf("Don't know the type, value: %v\n", t)
		}
	}

	whatAmI(false)
	whatAmI(1)
	whatAmI("hello")
}
