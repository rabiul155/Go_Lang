package main

import "fmt"

//Global variable
var x int = 100

//standard function. named function which has name
func add (a int , b int){ // here a and b are called parameter
	sum := a+b
	fmt.Println("This is normal function")
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

	// Main function print 500 because global variable x reassign from init function
	fmt.Println("print global variable change after init fn", x)

	// Normal function
	add(10, 15) // here 10 and 20 called parameter 
	

	//Anonymous function
	func (a, b int)  {
		sum := a+b
			fmt.Println("This is anonymous function")
		fmt.Println(sum)
	}(5,6) // This called Immediately Invoked Function Expression (IIFE)


	// anonymous function assign in a variable like arrow function in js
	add := func (a int , b int)  {
		sum := a+b
		fmt.Println("This is anonymous function assign in a variable")
		fmt.Println(sum)
	}
	add(200, 300)


	higherOrderFunction()

}

