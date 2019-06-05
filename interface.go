package main

import (
	"fmt"
	"math"
)

type Circle struct{

	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Square struct{
	Length float64
}

func (s *Square) Area() float64{
	return s.Length * s.Length	
}

func sumArea(shapes []Shape) float64{

	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()

	}
	return total
}

type Shape interface {
	Area() float64
}

func main(){
	
	s := &Circle{2.0}
	fmt.Println(s.Area())

	c := &Square{3.0}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	
	sa := sumArea(shapes)
	
	fmt.Println(sa)

}
