package main

import (

	"fmt"
)

func main(){

	loops := []string{"a", "bb", "ccc"}

	fmt.Printf("%v (type is %T )\n", loops, loops)

	fmt.Println(len(loops))

	fmt.Println(loops[1])


	for i, name := range loops {
		
		fmt.Println(i, name)	
	
	}

	loops = append(loops, "xyz")

	for _, k := range loops {
		fmt.Println(k)
	}

}
