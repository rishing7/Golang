package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	
	Area() float64

	Perim() float64

}

type Circle struct {

	Radius float64

}

func (c *Circle) Area() float64 {

	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perim() float64 {
	
	return 2 * math.Pi * c.Radius

}

type Rectangle struct {

	Length float64
	Width float64

}

func (r *Rectangle) Area() float64{

	return r.Length * r.Width
}

func (r *Rectangle) Perim() float64{

	return 2 * r.Length + 2 * r.Width
}

func Measure(g Geometry) {
	
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main(){
	
	c := &Circle{1.0}
	gc := Geometry(c)
	Measure(gc)

	r := &Rectangle{2.0, 3.0}

	gr := Geometry(r)
	Measure(gr)

}
