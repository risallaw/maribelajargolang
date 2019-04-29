package main

import "fmt"

func main() {
	fmt.Print("Nilai ujian: ")

	var nilai int
	_, err := fmt.Scanf("%d\n", &nilai)
	if err != nil {
		panic(err)
	}

	if nilai >= 60 {
		fmt.Println("Selamat Anda lulus.")
	} else {
		fmt.Println("Maaf Anda tidak lulus.")
	}
}
