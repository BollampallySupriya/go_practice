package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006")) // always use 01-02-2006 to format it in dd-mm-yy format.
	fmt.Println(presentTime.Format("01-02-2006 Monday")) // always use Monday to format in that way
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.June, 17, 13, 13, 13, 0, time.UTC)
	fmt.Println(createdDate)
}