package main

import "fmt"

func main() {
	var nilai int
	nilai = 90
	switch nilai {
	case 100:
		fmt.Println("Bagus Banget!")
		break
	case 90:
		fmt.Println("Bagus Aja")
		break
	case 80:
		fmt.Println("Bagus Lah")
		break
	default:
		fmt.Println("Belajar lagi Lah")
	}
}
