package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukan nilai (0-10): ")
	fmt.Scanf("%d\n", &nilai)

	switch nilai {
	case 0:
		fmt.Println("Nol")
	case 1, 3, 5, 7, 9:
		fmt.Println("Ganjil")
	case 2, 4, 6, 8, 10:
		fmt.Println("Genap")
	default:
		fmt.Println("Masukan tidak sesuai.")
	}
}
