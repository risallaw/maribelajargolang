package main

import "fmt"

func main() {
	nama := [5]string{"Andi", "Budi", "Cinta", "Dedi", "Edo"}
	for i, n := range nama {
		fmt.Printf("%d - %s\n", i, n)
	}

	text := "hello"
	for _, t := range text {
		fmt.Printf("%s.", string(t))
	}
	fmt.Println()
}
