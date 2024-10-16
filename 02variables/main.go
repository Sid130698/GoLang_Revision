package main

import "fmt"

func main() {
	var username string = "Siddharth"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	var smallInt uint8 = 255;
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type: %T \n", smallInt)
    var largeInt int64 = 9223372036854775807;
	fmt.Println(largeInt)
	fmt.Printf("Variable is of type: %T \n", largeInt)
	var smallFloat float32 = 255.4552343234;
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)


	walrusOperator := "this is walrus operator";
	fmt.Println(walrusOperator)

	var withOutType = "this is without type";
	fmt.Println(withOutType)

	
}
