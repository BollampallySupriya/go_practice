package main

import "fmt"

func main() {
	fmt.Println("Control Flow!!!")

	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10{
		result = "Watch Out"
	} else {
		result = "Exactly 10 Login Count"
	}
	fmt.Println(result)

	if num := 3 ; num < 3 {
		fmt.Println("Less than 3")
	} else {
		fmt.Println("Greater than or equal to 3")
	}
}