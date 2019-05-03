package main

import "fmt"

func main() {
	data := make(map[string]int)

loop:
	for {
		fmt.Println("1. Tambah/ubah data.")
		fmt.Println("2. Hapus data.")
		fmt.Println("3. Selesai.")

		var pilihan string
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		fmt.Println("")

		switch pilihan {
		default:
			fmt.Println("Pilihan tidak tersedia.")
		case "1":
			var key string
			var value int
			fmt.Println("Masukan key dan value (contoh: Apel 5).")
			fmt.Scanf("%s %d\n", &key, &value)
			data[key] = value
		case "2":
			var key string
			fmt.Print("Masukan key yang ingin dihapus: ")
			fmt.Scanln(&key)
			_, ok := data[key]
			if !ok {
				fmt.Println("Data tidak ditemukan.")
			}
			delete(data, key)
		case "3":
			break loop
		}

		fmt.Println()
		fmt.Printf("data = %v\n\n", data)
	}
}
