package main

import "fmt"

func main() {
	var deret = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(deret)

	// [2 3 4 5 6 7 8 9]
	fmt.Println(deret[2:])

	// [0 1 2 3 4 5]
	fmt.Println(deret[:6])

	// [2 3 4 5]
	fmt.Println(deret[2:6])
}
