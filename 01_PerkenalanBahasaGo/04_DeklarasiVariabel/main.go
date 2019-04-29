package main

import "fmt"

var nama = "Andi"

func main() {
	var umur int
	umur = 12

	berat := 26.5

	fmt.Printf("%v %T\n", nama, nama)
	fmt.Printf("%v %T\n", umur, umur)
	fmt.Printf("%v %T\n", berat, berat)
}
