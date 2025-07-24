package main

import "fmt"

const  p = 100 //declare const 
var a int = 10 //declare global variable

func outer () func () {
	inc :=0

	inner :=  func () {
		inc++
		fmt.Println(inc)
	}
	return inner
}

func main() {

 counter1 := outer()
 counter1()
 counter1()
 counter2 := outer()
 counter2()
 counter2()
	
}

func init (){
	fmt.Println("Introduction to closure")
}