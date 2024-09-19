package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	if day == "Monday"{
	// 		continue
	// 	}
	// 	fmt.Println(index, day)
	// 	if day == "Friday" {
	// 		break
	// 	}
	// }

	rogueValue := 1
	for rogueValue < 10 {
		if rogueValue == 8 {
			goto lco
		}
		if rogueValue == 7 {
			break
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println(rogueValue)
		rogueValue++
		
	}
	lco:
		fmt.Println("Using goto statements")
}