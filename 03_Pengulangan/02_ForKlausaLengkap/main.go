package main

import "fmt"

func main() {
	var n int
	fmt.Print("N = ")
	fmt.Scanf("%d\n", &n)

	faktorial := 1
	for i := 2; i <= n; i++ {
		faktorial *= i
	}
	fmt.Printf("Faktorial(N) = %v\n", faktorial)
}
