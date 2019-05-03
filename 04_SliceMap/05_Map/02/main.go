package main

import "fmt"

func main() {
	id2en := map[string]string{
		"senin":  "monday",
		"selasa": "tuesday",
		"rabu":   "wednesday",
		"kamis":  "thursday",
		"jumat":  "friday",
		"sabtu":  "saturday",
		"minggu": "sunday",
	}

	var hari string
	fmt.Print("Masukan nama hari (indonesia): ")
	fmt.Scanln(&hari)

	fmt.Printf("Nama hari dalam bahasa inggris: %s\n", id2en[hari])
}
