package main

import "fmt"

func main() {
	fmt.Println("About Methods in GO!!!")
	user := User{"Supriya", "supriya@gmail.com", true, 22}
	fmt.Printf("User object is %+v", user)
	userName := user.getUserName()
	fmt.Println(userName)
	user.changeEmail("priya@gmail.com")
	fmt.Printf("User object is %+v", user)
}

type User struct {
	Name string
	Email string 
	Status bool 
	Age int 
}

func (u User) getUserName() string{
	return u.Name
}

func (u User) changeEmail(email string) { // this does not modify the original object as the copy of object is passed here not the reference.
	u.Email = email
	fmt.Println("Email changed successfully and the new email is: ", u.Email)
}

// func (*u) permanentlyChangeEmail(email string) {
// 	*u.Email = email
// }