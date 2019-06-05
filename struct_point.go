package main

import (
	"fmt"
)
type Point struct{
	X int
	Y int
}


func (p *Point) Move(dx int, dy int){
	p.X += dx
	p.Y += dy

}

func main(){
	
	p := Point{
		X: 12,
		Y: 21,		
	}
	fmt.Printf("%+v\n",p)
	p.Move(2, 3)
	fmt.Printf("%+v\n",p)

}
