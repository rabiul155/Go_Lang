package main

import (
	"fmt"
	"mathlib/mathlib"
)

//global scope
var a int = 10
var b int = 100

func add(x int, y int){
	sum :=x+y
	
	fmt.Println(sum)
}

func variableStore (){
	var m int = 20
	var n int = 30

	fmt.Println(m,n)
	fmt.Println(a,b) // global variable can be access from here
}

func main (){

	p:=40
	q:=50

	add(p,q) // work with local variable

	add(a,b)// work with global variable

	// add(m,n) // does not work m and n does not exist in local or global variable
	variableStore()

	// sumVar(p,q) // get the function from same directory but another file 



	fmt.Println("Showing mathlib sum")
	mathlib.Math(p,q)

}