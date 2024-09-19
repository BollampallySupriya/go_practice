package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate our app between 1 to 5:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // read till new line

	// numRating, err := strconv.ParseFloat(input, 64)	// Error Occurred strconv.ParseFloat: parsing "4\r\n": invalid syntax (As it also reads new line)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error Occurred", err)
	} else {
		fmt.Println(numRating + 1)
	}
}
