package main

import "fmt"

func main() {
	fmt.Println("All about Arrays!!!")

	var fruitList [4]string;
	fruitList[0] = "banana"
	fruitList[1] = "apple"
	fruitList[3] = "pomegranate"
	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var vegList = [3]string {"potato", "tomato", "bendi"}
	fmt.Println(vegList)
}