package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	// Variables
	var name string = "Hrushi"
	fmt.Printf("Hey there! I am %s\n", name)

	age := 22
	fmt.Printf("I am %d years old\n", age)

	var city string = "Pune"

	country := "India"
	fmt.Printf("I am from %s, %s\n", city, country)

	// Zero Values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBoolean bool

	fmt.Printf("Default Value\n defaultInt: %d \n defaultFloat: %f \n defaultString: %s \n defaultBoolean: %t \n", defaultInt, defaultFloat, defaultString, defaultBoolean)

	// Constants and Enums
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

	// Conditional Statement
	if age >= 18 {
		fmt.Printf("You are %d old so you are an adult\n", age)
	} else if age <= 13 {
		fmt.Printf("You are %d old so you are a teenager\n", age)
	} else {
		fmt.Printf("You are %d old so you are a child\n", age)
	}

	day := "Sunday"

	switch day {
	case "Monday", "Tuesday":
		fmt.Println("Start of the week")
	case "Wednesday", "Thursday", "Friday":
		fmt.Println("Mid week")
	default:
		fmt.Println("Weekend")
	}

	// Loops
	for i := 0; i < 5; i++ {
		fmt.Printf("For Loop : Value of i = %d\n", i)
	}

	counter := 0
	for counter < 3 {
		fmt.Printf("Simulating while loop using for with conditional statement : counter = %d\n", counter)
		counter++
	}

	iterations := 0

	for {
		if iterations > 3 {
			break
		}
		fmt.Printf("Simulating infinite while loop using for with base condition: itterations =  %d\n", iterations)
		iterations++
	}

	// Arrays

	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("Array of numbers = %v\n", numbers)
	fmt.Printf("Length of Array of numbers = %d\n", len(numbers))
	fmt.Printf("First Element in Array of numbers = %d\n", numbers[0])

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("Matrix of numbers = %v\n", matrix)

	// Slices
	allNumbers := numbers[:]
	firstThree := numbers[0:3]
	fmt.Printf("allNumbers slice = %v\n", allNumbers)
	fmt.Printf("Append numbers 60 to all numbers = %v\n", append(allNumbers, 60))
	fmt.Printf("Slice of first 3 elements of numbers = %v\n", firstThree)

	fruits := []string{"Mango", "Apple"}
	fmt.Printf("Slice of Fruits = %v\n", fruits)
	fmt.Printf("Append orange, pineapple, papaya to Slice of Fruits = %v\n", append(fruits, "Orange", "Pineapple", "Papaya"))

	fmt.Println("Itterating numbers with range")
	for index, value := range numbers {
		fmt.Printf("index %d and value %d\n", index, value)
	}
}

func add(a int, b int) int {
	return a + b
}

func calcSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
