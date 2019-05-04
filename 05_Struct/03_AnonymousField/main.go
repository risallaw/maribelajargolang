package main

import "fmt"

type data struct {
	string
	int
}

func main() {
	var d data
	fmt.Print("Masukan nama dan umur: ")
	fmt.Scanf("%s %d\n", &d.string, &d.int)

	fmt.Println()
	fmt.Printf("Nama: %s\nUmur: %d\n", d.string, d.int)
}
