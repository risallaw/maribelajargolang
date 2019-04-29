package main

import "fmt"

func main() {
	var nilai float32
	fmt.Print("Nilai (0-100): ")
	fmt.Scanf("%f\n", &nilai)

	if !(nilai >= 0 && nilai <= 100) {
		fmt.Println("Masukan Anda salah.")
	}
}
