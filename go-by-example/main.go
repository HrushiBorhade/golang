package main

import (
	"fmt"

	"github.com/HrushiBorhade/golang/go-by-example/sections/channels"
	"github.com/HrushiBorhade/golang/go-by-example/sections/customerror"
	"github.com/HrushiBorhade/golang/go-by-example/sections/goroutines"
)

func main() {
	fmt.Println("Go By Example!")
	fmt.Println("-------------- Errors -------------")
	customerror.Execute()
	fmt.Println("-------------- Goroutines -------------")
	goroutines.Execute()
	fmt.Println("---------------- Channels ---------------")
	channels.Execute()
	fmt.Println("-----------------------------------")
}
