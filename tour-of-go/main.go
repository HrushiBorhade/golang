package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

var (
	ToBe    bool       = false
	MaxInt  uint64     = 1<<64 - 1
	Complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println("My fav number is", rand.Intn(10))
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	var x, y, z = true, false, "no!"
	fmt.Println(x, y, z, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Complex, Complex)

	const World = "世界"
	fmt.Println("Hello", World)
}
