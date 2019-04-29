package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukan nilai: ")
	fmt.Scanf("%d\n", &nilai)

	switch {
	case nilai < 0:
		fmt.Println("Negatif")
	case nilai == 0:
		fmt.Println("Nol")
	case nilai%2 == 1:
		fmt.Println("Ganjil")
	default:
		fmt.Println("Genap")
	}
}
