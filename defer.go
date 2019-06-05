package main
import (
	"fmt"
)

func cleanup(s string){
	
	fmt.Printf("%v Cleaning...", s)

}

func worker(){
	defer cleanup("A")
	
	fmt.Println("In worker...")
}
func main(){

	worker()
}
