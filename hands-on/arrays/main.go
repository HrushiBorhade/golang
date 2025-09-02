package main 

import "fmt"

func main() {
	var a [5]int
	fmt.Printf("empty array a:%v \n",a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:",a[4])

	b := [5]int{1,2,3,4,5}
	fmt.Println("Declarative: ",b)
	
	c := [...]int{9,8,7,6}
	fmt.Printf("Spread operator for array %v with length %d \n" , c, len(c))

	d := [...]int{1,3:4,5,6}
	fmt.Println("indexed declaration", d)


	var twoD [2][3]int;

	for i:= range 2{
		for j:= range 3{
			twoD[i][j] = i+j
		}
	}

	fmt.Println("TWO DIMENSIONAL ARRAY",twoD)
}

