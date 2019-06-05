package main

import "fmt"

func main(){
	
	var i int;
	//i = 20	Default value is 0
	
	var f float64;
	f = 20.4;
	
	var c byte;
	c = 'a'
	

	fmt.Println(i)
	fmt.Printf("Type of i is %T\n", i)

	fmt.Println(f)
	fmt.Printf("Type of f is %T\n", f)

	fmt.Println(c)
	fmt.Printf("Type of c is %T\n", c)	
	
	d := 100;	
	
	fmt.Println(d)
	fmt.Printf("Type of d is %T\n", d)

	
	var x,y,z = 9, 20.4, "rishi"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	m := 20.2
	fmt.Println(m)
	
	const length int = 4
	const width int = 5
	
	var area int
	
	area = length * width	
	
	fmt.Println(area)

	fmt.Println("Hello World!")
}
