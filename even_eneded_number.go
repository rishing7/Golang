package main

import (
	"fmt"
)

func main(){

	x := 123121

	s := fmt.Sprintf("%d", x)

	//	fmt.Println(s)
	
	if s[0] == s[len(s)-1] {
		fmt.Println("This is even ended number")
	} else {
	
		fmt.Println("This is not even ended number")
	}

}
