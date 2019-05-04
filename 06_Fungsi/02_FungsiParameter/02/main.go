package main

import "fmt"

func luasSegitiga(alas, tinggi float64) {
	luas := 0.5 * alas * tinggi
	fmt.Printf("Luas segitiga: %v\n", luas)
}

func main() {
	luasSegitiga(4, 5)
}
