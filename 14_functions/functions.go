package main

import "fmt"

func main() {
	fmt.Println("All about functions!!!")
	result := add(3, 4)
	fmt.Println(result)
	fmt.Println(greet())
	msg, res := adder(1, 2, 3, 4, 5)
	fmt.Println(msg, res)
}

func add(a int, b int) int {
	var sum = a + b 
	return sum
}

func adder(args ...int) (string, int) {
	var sum = 0
	for _, value := range args {
		sum = sum + value
	}
	return "The result is: ", sum
}

func greet() string{
	return "Namaste Folks!!"
}