package main 

import (
	"fmt"
	"slices"
)

func main() {
	 var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0, cap(s))

	s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])


	s = append(s,"d" )
	s =append(s,"e", "f")
	fmt.Println("s: ", s, len(s), cap(s))


	d :=  make([]string, len(s))

	copy(d,s)

	fmt.Println("d",d)


	t1 := []string{"x","y","z"}

	t2 := []string{"x", "y","z"}

	if slices.Equal(t1,t2) {
		fmt.Println("t1 and t2 are equal")
	}
}

