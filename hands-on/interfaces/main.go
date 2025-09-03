package main 


import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
}

type rect struct {
	width float64
	height float64
}

type circle struct{
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi* c.radius*c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area: ",g.area())
	fmt.Println("perim: ",g.perim())
}

func detectShape(g geometry) {
	switch t := g.(type) {
		case circle:  {
			fmt.Println("it's a circle: ",t)
    	}
	case rect: {
		fmt.Println("it's a rectangle: ",t)
	}
default:
	fmt.Println("No fucking idea what shape this is",t)
	}
}

func main() {
	r :=rect{width:10, height:20}
	c :=circle{radius:10}

	measure(r)
	measure(c)

	detectShape(r)
	detectShape(c)
}
