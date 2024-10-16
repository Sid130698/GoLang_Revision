package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza...")
	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Printf("Type of Rating %T ", input)

}
