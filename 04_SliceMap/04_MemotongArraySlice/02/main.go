package main

import (
	"fmt"
	"os"
)

func main() {
	array := []int{}
	for {
		fmt.Println("1. Tambah data.")
		fmt.Println("2. Hapus data.")
		fmt.Println("3. Selesai.")

		var pilihan string
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		fmt.Println()

		switch pilihan {
		case "1":
			var nilai int
			fmt.Print("Nilai: ")
			fmt.Scanln(&nilai)
			array = append(array, nilai)
			fmt.Printf("Array: %v\n\n", array)
		case "2":
			var indeks int
			fmt.Print("Indeks: ")
			fmt.Scanln(&indeks)
			if indeks >= 0 && indeks < len(array) {
				array = append(array[:indeks], array[indeks+1:]...)
			} else {
				fmt.Println("Indeks tidak ditemukan.")
			}
			fmt.Printf("Array: %v\n\n", array)
		case "3":
			os.Exit(0)
		}
	}
}
