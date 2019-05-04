package main

import "fmt"

func luasKelilingPersegiPanjang(panjang, lebar float64) (float64, float64) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

func main() {
	var p, l float64
	fmt.Print("Masukan panjang dan lebar: ")
	fmt.Scanf("%f %f\n", &p, &l)

	luas, keliling := luasKelilingPersegiPanjang(p, l)
	fmt.Printf("Luas: %v\nKeliling: %v\n", luas, keliling)
}
