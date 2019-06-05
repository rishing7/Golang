package main

import (

	"fmt"
)

func add(a int, b int) int {

	return a + b
}

func div_mod(a int, b int) (int, int){

	return a/b, a%b

}

func doubleAt(arr []int, i int){

	arr[i] *= 2
}

func main(){

	val := add(2, 4)
	fmt.Println(val)

	div, mod := div_mod(5, 2)
	fmt.Println(div, mod)

	arr := []int{1, 5, 3, 6, 2}
	fmt.Println(arr)
	doubleAt(arr, 3)
	fmt.Println(arr)

}
