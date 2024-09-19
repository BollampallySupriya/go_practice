package main

import "fmt"

func main() {
	defer fmt.Println("First Defer Statement!!")
	defer fmt.Println("Second Defer Statement!!")
	fmt.Println("DEFER Statement!!!")
	defer myDefer()
	defer fmt.Println("Third Defer Statement!!")
	myDefer()
}

func myDefer() {
	defer fmt.Printf("\n")
	for i:=0 ; i<5; i++ {
		defer fmt.Printf("%v\t", i)
	}
}