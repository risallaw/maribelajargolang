package main

import "fmt"

func main() {
	s := struct {
		alas, tinggi float64
	}{
		alas:   4,
		tinggi: 5,
	}

	fmt.Printf("Luas = %v\n", 0.5*s.alas*s.tinggi)
}
