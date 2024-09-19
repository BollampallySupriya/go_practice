package main

import "fmt"

func main() {
	fmt.Println("ALl about maps!!!")
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"
	// fmt.Println("List of all languages: ", languages)
	// fmt.Println("JS shorts for : ", languages["JS"])

	delete(languages, "RB")
	// fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang
	// similar to for each loop in Javascript
	for key, value := range languages {
		fmt.Printf("Key: %v , Value: %v\n", key, value)
	}
}