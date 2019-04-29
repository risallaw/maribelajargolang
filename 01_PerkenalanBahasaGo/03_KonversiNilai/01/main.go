package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sisiStr string
	fmt.Print("Sisi: ")
	fmt.Scanln(&sisiStr)

	sisi, err := strconv.ParseFloat(sisiStr, 64)
	if err != nil {
		panic(err)
	}

	luas := sisi * sisi
	fmt.Printf("Luas persegi %v\n", luas)
}
