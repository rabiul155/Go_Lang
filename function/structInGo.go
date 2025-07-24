package main

import "fmt"

type User struct { // like type in js
	name string
	age  int
}

func structInGo() {

	var user1 User // variable declare

	user1 = User{ // value assign 
		name: "Rahim",
		age:  12,
	}

	fmt.Println(user1)

	user2 := User{ // variable declare and assign
		name: "Karim",
		age:  14,
	}

	fmt.Println(user2)

}
