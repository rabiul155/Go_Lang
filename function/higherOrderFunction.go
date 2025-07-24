package main

import "fmt"


func addNumber(a int , b int){
	sum := a+b
	fmt.Println(sum)
}

func processFn (x int , y int , fn func(p int , q int)){
	fmt.Println("I am higher order function i can take fn as parameter")
	fn(x,y)
}

func returnAnotherFn ()func (a int , b int){
	fmt.Println("I am higher order function i can return another fn")
	return addNumber
}


func higherOrderFunction (){
	processFn(4,5 ,addNumber)

	// Receive returned function and call 
	getFn := returnAnotherFn()
	getFn(5,6)

	// call directly returned function
	returnAnotherFn()(4,5)
}