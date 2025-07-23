package main

import "fmt"

//Global variable
var x int = 100

//standard function. named function which has name
func add (a int , b int){
	sum := a+b
	fmt.Println(sum)
}


// init function name always init and it does not callable form another function
// computer call init function automatically and it calls before the main function 
func init(){
	fmt.Println("I am init function")
	fmt.Println("print global variable", x)
	// Change global variable 
	x=500

}


func main (){
	fmt.Println("Function types")
	add(10, 15) 
	// Main function print 500 because global variable x reassign from init function
	fmt.Println("print global variable", x)

	//Anonymous function
	func (a, b int)  {
		sum := a+b
		fmt.Println(sum)
	}(5,6) // This called Immediately Invoked Function Expression (IIFE)

}

