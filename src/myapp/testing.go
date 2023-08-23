package main

import "fmt"

func main() {
	nilai := 90

	switch nilai {
	case 100:
		fmt.Println("Mantep")
	case 90:
		fmt.Println("Sabi")
	case 80:
		fmt.Println("Ok Lah")
	case 70:
		fmt.Println("Belajar lah")
	default:
		fmt.Println("Belajar lagi")
	}
}
