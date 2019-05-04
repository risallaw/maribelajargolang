package main

import "fmt"

type segitiga struct {
	alas   float64
	tinggi float64
}

func main() {
	var s segitiga

	fmt.Print("Alas: ")
	fmt.Scanf("%f\n", &s.alas)

	fmt.Print("Tinggi: ")
	fmt.Scanf("%f\n", &s.tinggi)

	fmt.Printf("0.5 x %v x %v = %v\n",
		s.alas, s.tinggi, 0.5*s.alas*s.tinggi)
}
