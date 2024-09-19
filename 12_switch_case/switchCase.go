package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("About Switch Case!!!")
	rand.Seed(time.Now().UnixNano())
	var diceNumber = rand.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println("Open your game by moving pawn to 1")
	case 2:
		fmt.Println("Move your pawn by 2")
	case 3:
		fmt.Println("Move your pawn by 3")
		fallthrough
	case 4:
		fmt.Println("Move your pawn by 4")
	case 5: 
		fmt.Println("Move your pawn by 5")
	case 6:
		fmt.Println("Move your pawn by 6 and Roll the dice again!")
	default:
		fmt.Println("Invalid Choice")
	}
}