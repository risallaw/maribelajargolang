package main

import "fmt"

func main() {
	var nama string
	fmt.Print("Nama: ")
	fmt.Scanln(&nama)

	var jumlah int
	fmt.Print("Jumlah pengulangan: ")
	fmt.Scanf("%d\n", &jumlah)

	var i int
	for i < jumlah {
		fmt.Println(nama)
		i++
	}
}
