package main 

import "fmt"

func main() {
	fmt.Println("SIMPLE WHILE LOOP KINDA ITTERATION")
	i := 0
	for i < 3 {
		fmt.Println("i = ", i)
		i+=1
	}

	fmt.Println("CLASSIC CONDITIONS LOOP")
	for j:=0 ; j<3 ; j++ {
		fmt.Println("j =",j)
	}

	fmt.Println("RANGE")
	for i:= range 3 {
		fmt.Println("i range = ", i)
	}

	fmt.Println("LOOP")
	for {
		fmt.Println("loop ends here we break now")
		break
	}

	fmt.Println("CONTINUE")
	for i:=range 6 {
		if i%2==0 {
			continue
		}
		fmt.Println(i)
	}

}
