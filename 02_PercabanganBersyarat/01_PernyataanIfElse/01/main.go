package main

import "fmt"

func main() {
	fmt.Print("Nilai ujian: ")

	var nilai int
	fmt.Scanf("%d\n", &nilai)
	if nilai >= 60 {
		fmt.Println("Selamat Anda lulus.")
	}
}
