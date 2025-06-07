package main
import "fmt"

// Global Scope Variable
const (
	a = 10
	b = 20
)

//  Add Tow Number
func add(a,b int) int{
	return a+b
}

func main(){

	const x int= 90
	var data = add(x, a)
	fmt.Println(data)
}