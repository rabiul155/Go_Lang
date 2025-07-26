package main

import "fmt"

type User struct{
	name string
	age int
}

// Normal function
func printUserDetails (user User){
	fmt.Println("User name", user.name)
	fmt.Println("User age", user.age)
}


// Receiver function
func (user User) printDetails (){
	fmt.Println("User name", user.name)
	fmt.Println("User age", user.age)
}

func (user User) callUser (age int){
		fmt.Println("User name", user.name)
		fmt.Println("User age", user.age + age)
}


func main() {
	fmt.Println("Introduction to receiver function")

	var user1 User

	user1 = User{
		name : "Abul",
		age :13,
	}

	//Normal fn call 
	printUserDetails(user1)

	//Receiver fn call
	user1.printDetails()

	user2 :=User{
		name : "Kasem",
		age : 23,
	}

	printUserDetails(user2)

	user2.callUser(12)


}