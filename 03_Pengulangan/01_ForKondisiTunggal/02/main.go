package main

import "fmt"

func main() {
	var total int

	for {
		var nilai int
		fmt.Print("Nilai (integer positif): ")
		_, err := fmt.Scanf("%d\n", &nilai)
		if err != nil || nilai < 0 {
			break
		} else {
			total += nilai
		}
	}

	fmt.Printf("Total: %d\n", total)
}
