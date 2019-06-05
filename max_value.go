package main

import (
	
	"fmt"
)

func main(){

	arr := []int{2, 4, 7, 1, 3, 9, 5}
	max := arr[0]
	
	for i := range arr[1:] {
		
		if arr[i]>max{
			
			max = arr[i]
		}
	
	}
	
	fmt.Println(max)

}
