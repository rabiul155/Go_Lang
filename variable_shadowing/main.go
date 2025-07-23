package main

import "fmt"

var a int = 10

func main (){
	fmt.Println("Variable shadowing")

	age := 18

	if age >=18 {

		// this is called variable shadowing. assign a variable multiple time in multiple scope
		a := 50

		fmt.Println("a is ", a) // local scope variable a will be 50
	}

	fmt.Println("Global a is ", a) // global scope variable a will be 10
	
}