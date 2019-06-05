package main
import(
	
	"fmt"
)
func sum(nums... int){
	
	fmt.Print(nums, " ")
	
	total := 0

	for _, i := range nums {

		total += i
	}
	
	fmt.Print(total,"\n")

}

func main(){
	
	sum(1, 2, 3)
}
