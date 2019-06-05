package main

import (
     
      "fmt"
)

func main(){
     
     /*var x int
     var y int
     
     x = 1     
     y = 2 
	
	OR  */

     x := 2.0
     y := 3.0

     fmt.Printf("Value of x=%v , type is %T\n", x, x)
     fmt.Printf("Value of y=%v , type is %T\n", y, y)
	
     
     mean := (x+y)/2

     fmt.Printf("Mean is %v , type is %T \n", mean, mean) 
}
