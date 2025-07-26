package main

import "fmt"

func main() {
	fmt.Println("Introduction to array")

	// array declare in go
	var arr[2] int
	arr[0] = 4
	arr[1]  = 5

	arr2 := [2]int{5,6}

	fmt.Println(arr)
	fmt.Println(arr2)
 }