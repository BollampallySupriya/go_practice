package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with Files!!!")
	content := "This needs to go into a file - Learncodeonline.com"
	file, err := os.Create("./myfile.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	print("Content written to file and the length of the file is ", length)
	defer file.Close() 	// this helps to close the file after the completion of program execution.
	readFile(file.Name())
	// fmt.Println(file.Name())
}

func readFile(file_path string) {
	byte_data, err := os.ReadFile(file_path)
	checkNilError(err)
	fmt.Println("Content of the file is:\n", byte_data)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)		// shuts down the program and shows the error.
	}
}