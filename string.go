package main

import (
	"fmt"
)

func main(){
	
	book := "The color of magic"
	fmt.Println(book)
	
	fmt.Println(len(book))
	
	fmt.Printf("book[0] = %v ( type %T )\n", book[0], book[0])

	// String in go is immutable
	// book[0] = 119
	
	/* 	Slicing		*/
	fmt.Println(book[4:11])
	fmt.Println(book[4:])
	fmt.Println(book[:11])

}
