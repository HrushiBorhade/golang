package main 

import "fmt"

func causePanic(){
	panic("aaagghhhhh!!!")
}

func main() {
	fmt.Println("Before panic caused")

	defer func() {
		if r:= recover(); r!=nil {
			fmt.Println("Recovered, ERROR: ",r)
		}
	}()

	causePanic()

	fmt.Println("After panic is caused")
}
