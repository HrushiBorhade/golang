package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	//Variables
	var name string = "Hrushi"
	fmt.Printf("Hey there! I am %s\n", name)

	age := 27
	fmt.Printf("I am %d years old\n", age)

	var city string = "Pune"

	country := "India"
	fmt.Printf("I am from %s, %s\n", city, country)

	//Zero Values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBoolean bool

	fmt.Printf("Default Value\n defaultInt: %d \n defaultFloat: %f \n defaultString: %s \n defaultBoolean: %t \n", defaultInt, defaultFloat, defaultString, defaultBoolean)

	//Constants and Enums
	const PI = 3.14
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
	)
	fmt.Printf("Mon - %d\nTue - %d\nWed - %d\n", Monday, Tuesday, Wednesday)

	fmt.Println("Addition Result of 3 and 4 = ", add(3, 4))

	sum, product := calcSumAndProduct(4, 5)
	fmt.Printf("Sum of 4 and 5 = %d and product of 4 and 5 = %d\n", sum, product)
}

func add(a int, b int) int {
	return a + b
}

func calcSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
