package main

import "fmt"

func luasSegitiga(alas, tinggi float64) float64 {
	luas := 0.5 * alas * tinggi
	return luas
}

func main() {
	alas, tinggi := 4.0, 5.0

	fmt.Printf("LuasSegitiga(%v, %v) = %v\n",
		alas, tinggi, luasSegitiga(alas, tinggi))
}
