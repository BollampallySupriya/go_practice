package main

import "fmt"

func main() {
	fmt.Println("About Struct in GO!!!")
	// no inheritance in golang , no super or parent

	user := User{"Supriya", "supriya@gmail.com", true, 22}
	fmt.Println(user)
	fmt.Printf("User is %+v\n", user) // gives key value pairs
	fmt.Printf("Type of user is %T\n", user)	
	fmt.Printf("User name is %v\n", user.Name)
}

type User struct {
	Name string
	Email string 
	Status bool 
	Age int 
}