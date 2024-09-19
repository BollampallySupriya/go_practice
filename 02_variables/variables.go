package main
import "fmt"

func main(){
	fmt.Println("Variables")
	var username string = "Supriya"
	fmt.Printf("Username is of type %T \n Username is %v", username, username)

	var isUser bool = false
	fmt.Printf("Username is of type %T \n Username is %v", isUser, isUser)

	var num1 uint8 = 255 // 266 (cannot accept it as max size is 255 for 8 bits. Unsigned accepts values from 0)
	fmt.Printf("Username is of type %T \n Username is %v", num1, num1)

	var num2 uint16 = 16578 // max is 16 bits and it is unsigned that is accepts values from 0
	fmt.Printf("Username is of type %T \n Username is %v", num2, num2)

	var num3 int8 = 121 // max is 16 bits and it is signed that is accepts values from -126 to 127
	fmt.Printf("Username is of type %T \n Username is %v", num3, num3)

	var num4 int16 = 4567 // max is 16 bits and it is signed that is accepts values from - to +
	fmt.Printf("Username is of type %T \n Username is %v", num4, num4)

	var num5 float32 = 4567.34345682392356239531  // gets 5 values after the decimal point
	fmt.Printf("Username is of type %T \n Username is %v", num5, num5)

	var num6 float64 = 456767839.3674258375981375893 // more precision after the decimal.
	fmt.Printf("Username is of type %T \n Username is %v", num6, num6)


	// implicit declaration (go automatically picks up suitable type based on value if not mentioned.)
	var company = "Google"
	fmt.Println(company)

	// no var style (only allowed inside the methods . not outside functions.)
	place:= "Hyderabad"
	fmt.Println(place)
}