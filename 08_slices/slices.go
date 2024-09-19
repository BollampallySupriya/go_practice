package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("All about Slices!!!")

	var fruitSlice = []string{"Apple", "Banana"} // we have to initialize and should not give length in order to make it slice.
	// fmt.Printf("Type of fruitSlice is %T \n", fruitSlice)
	fmt.Println("Length of fruitSlice is, ", len(fruitSlice))
	fmt.Println("Capactity of fruitSlice is, ", cap(fruitSlice))

	fruitSlice = append(fruitSlice, "PomeGranate")

	fmt.Println("Length of fruitSlice is: ", len(fruitSlice))
	fmt.Println("Capactity of fruitSlice is: ", cap(fruitSlice))

	fruitSlice = append(fruitSlice[1:])
	fmt.Println(fruitSlice)

	highScore := make([]int, 4)
	fmt.Printf("Type of highScore is: %T \n", highScore)
	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 456
	highScore[3] = 674
	// highScore[4] = 274 // this will give array out of bound exception
	highScore = append(highScore, 564, 679) // this will work as it automatically makes it as slice.
	sort.Ints(highScore) // this is sort the integer values in the slice
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))


	// how to remove value from slices based on index.
	var courses = []string{"reactjs", "javascript", "python", "go", "java"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
