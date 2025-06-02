package main
import "fmt"

func main(){
	const age int= 20
	// if age >= 18{
	// 	fmt.Println("You are eligible for Married")
	// } else if age <18{
	// 	fmt.Println("You are not eligible for Married but you can love someone")
	// }else {
	// 	fmt.Println("You are not eligible for Married")
	// }
	



	if age == 20 { // 18 or equals
		fmt.Println("You are Perfect eligible for Married")
	}else if age <= 18{
		fmt.Println("You are  eligible for Married")
	}else {
		fmt.Println("You are eligible for Married")
	}


	const x int = 30
	const gender string = "Male"
	//  And &&
	if x == 20 && gender == "Male"{
		fmt.Println("Your age is 20 and your gender is Male")
	}

	// Or ||
	if x == 30 || gender == "Male"{
		fmt.Println("Your age is 30 or your gender is Male")
	}

	// Not !
	const isAdmin bool = false
	if !isAdmin{
		fmt.Println("You are not admin")
	}

}