package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	//comma ok syntax or error ok
	input, _ := reader.ReadString('\n')	// will get value in input if it is success else in _(which is error)
	fmt.Println("Hello, Good Morning  ", input)
	fmt.Printf("Type of input is %T", input) // will be string always irrespective of your input type. (as we r reading string)
	
}