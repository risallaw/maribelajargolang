package main

import (
	"fmt"
	"strconv"
)

func main() {
	var count int
	daftarNilai := []float64{}
	for {
		var nilaiStr string
		fmt.Print("Masukan nilai (x - berhenti): ")
		fmt.Scanln(&nilaiStr)
		if nilaiStr == "x" {
			break
		}

		nilai, err := strconv.ParseFloat(nilaiStr, 64)
		if err != nil {
			fmt.Println("Masukan tidak sesuai, silahkan coba lagi.")
		} else {
			daftarNilai = append(daftarNilai, nilai)
			count++
		}
	}

	fmt.Println(daftarNilai)
	var total float64
	for _, n := range daftarNilai {
		total += n
	}

	rataRata := total / float64(count)
	fmt.Printf("\nRata-rata nilai = %f\n", rataRata)
}
