package main

import (
	"fmt"
	"strings"
)

func main() {
	var teks string
	fmt.Print("Masukan sebuah teks: ")
	fmt.Scanln(&teks)

	lowercase := strings.ToLower(teks)
	var vokal, konsonan int
	for _, t := range lowercase {
		switch string(t) {
		case "a", "i", "u", "e", "o":
			vokal++
		case "b", "c", "d", "f", "g", "h", "j",
			"k", "l", "m", "n", "p", "q", "r",
			"s", "t", "v", "w", "x", "y", "z":
			konsonan++
		}
	}

	fmt.Printf("Vokal: %d\n", vokal)
	fmt.Printf("Konsonan: %d\n", konsonan)
}
