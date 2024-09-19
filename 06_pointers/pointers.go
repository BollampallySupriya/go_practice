package main

import "fmt"

func main() {
	fmt.Println("All About Pointers!!")

	// var ptr *int; // this declares the ptr as the reference to the memory location. Default value is nil
	// fmt.Printf("Type of ptr is %T ", ptr)

	myNumber := 22
	var myPtr = &myNumber // referencing myNumber soo we use '&'
	fmt.Printf("Type of ptr is %T \n", myPtr)
	fmt.Println("Value of ptr is ", *myPtr)
	fmt.Println("Address of ptr is ", myPtr)
	fmt.Println("My number is ", myNumber)

	// changing the original value of myNumber in memory location
	*myPtr = *myPtr * 2 
	fmt.Println(" My number after changing is ", myNumber)
}