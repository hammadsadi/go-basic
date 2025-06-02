package main
import "fmt"

func main() {
	const x int = 4
	switch x {
	case 1:
		fmt.Println("One")
	case 2, 4:
		fmt.Println("Two or Four")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("None")
	}
}