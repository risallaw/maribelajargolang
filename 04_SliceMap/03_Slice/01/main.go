package main

import "fmt"

func main() {
	var nilai []int
	nilai = make([]int, 0, 2)

	for i := 0; i < 5; i++ {
		fmt.Printf("Len: %d, Cap: %d, Nilai: ", len(nilai), cap(nilai))

		var n int
		fmt.Scanf("%d\n", &n)
		nilai = append(nilai, n)
	}

	fmt.Printf("Nilai: %v\n", nilai)
}
