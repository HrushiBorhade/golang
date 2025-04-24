package main

import (
	"fmt"

	error_test "github.com/HrushiBorhade/golang/go-by-example/sections/error"
	"github.com/HrushiBorhade/golang/go-by-example/sections/goroutines"
)

func main() {
	fmt.Println("Go By Example!")
	fmt.Println("-------------- Errors -------------")
	error_test.Execute()
	fmt.Println("-------------- Goroutines -------------")
	goroutines.Execute()
	fmt.Println("-----------------------------------")
}
