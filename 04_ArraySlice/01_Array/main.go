package main

import "fmt"

func main() {
	var nilaiSiswa [5]int
	for i := 0; i < len(nilaiSiswa); i++ {
		fmt.Printf("Nilai siswa ke-%d: ", i+1)
		fmt.Scanf("%d\n", &nilaiSiswa[i])
	}

	fmt.Printf("Nilai siswa: %v\n", nilaiSiswa)

	var rataRataNilai float64
	for _, n := range nilaiSiswa {
		rataRataNilai += float64(n)
	}

	rataRataNilai /= float64(len(nilaiSiswa))
	fmt.Printf("Rata-rata nilai: %f\n", rataRataNilai)
}
