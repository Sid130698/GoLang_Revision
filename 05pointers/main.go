package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")
	// var ptr *int
	// fmt.Println("the value of pointer is ", ptr)
	myNumber := 24
	var ptr *int = &myNumber
	fmt.Println("value of ptr", *ptr)

}
