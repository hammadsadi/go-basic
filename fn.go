package main
import "fmt"

func add(x int, y int){
	var sum = x + y
	fmt.Println(sum)

}


func getNumbers (x int, y int) ( int, int) {
	sum := x + y
	diff := x - y
	return sum, diff
}

func main() {


	// add(1, 2)
	// add(5, 5)

	x, y := getNumbers(10, 5)
	fmt.Println(x)
	fmt.Println(y)

}