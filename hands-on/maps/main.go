package main 

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Printf("m :%v \n",m)

	delete(m, "k2")
    fmt.Println("map:", m)

	clear(m)
    fmt.Println("map:", m)

	 _, prs := m["k2"]
    fmt.Println("prs:", prs)

	  n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	 n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }
}
