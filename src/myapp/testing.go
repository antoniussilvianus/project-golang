package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a Number: ")
	var input float32
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	output := input * 2
	fmt.Println(output)
}
