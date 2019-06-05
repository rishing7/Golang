package main

import (

	"fmt"
)

func f(s string){
	
	for i:=0; i<3; i++ {
		fmt.Println(s, " : ", i)
	}

}

func main(){

	go f("goroutine")
	
	go func(msg string){
		fmt.Println(msg)
	}("going")
	
	f("direct")

	fmt.Scanln()

	fmt.Println("Done!!!")

}
