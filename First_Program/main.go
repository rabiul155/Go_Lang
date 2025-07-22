package main

import "fmt"

func helloWorld (){
	fmt.Println("Hello, world!")

	//declare variable 
	var a int = 10
	fmt.Println(a);

	//Variable declare sort-hand
	x := 100
	fmt.Println(x)

	//variable redeclare 
	m :=true
	m = false
	fmt.Println(m)
}

func ifelse (){
	a := 10

	if a > 10 {
		fmt.Println("Number is greater than 10")
	}else if a < 10 {
		fmt.Println("Number is less than 10")
	} else{
		fmt.Println("Number is 10")
	}

}

func switchCase (){
	a :=3

	switch a {
	case 1 :{
		fmt.Println("a is one")
		
	}
	case 2, 3 : {
		fmt.Println("a is either two or thee")
	}
	default : {
		fmt.Println(" a is nothing")
	}
	}
}

func addFunc (num1 int, num2 int){

	sum := num1 + num2

	fmt.Println(sum)
}

func returnFunc (num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func multipleReturnFunc (num1 int, num2 int) (int, int){
	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func sayHello (name  string){


	fmt.Println("say hello to", name)
}

func main(){
	helloWorld()
	ifelse()
	switchCase()
	a:=10
	b:=20
	addFunc(a,b)
	sum := returnFunc(a,b)
	fmt.Println(sum)
	sum2 , mul := multipleReturnFunc(a,b)
	fmt.Println(sum2, mul)
	sayHello("siyam")

}