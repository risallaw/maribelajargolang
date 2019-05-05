package main

import (
	"fmt"
	"math"
	"time"
)

func printChar(text string) {
	for _, c := range text {
		fmt.Printf("%v ", string(c))
		time.Sleep(10 * time.Millisecond)
	}
}

func printDeret(a, r, n0, n1 float64) {
	if n1 < n0 {
		n0, n1 = n1, n0
	}

	rn := math.Pow(r, n0)
	for n := n0; n <= n1; n++ {
		fmt.Printf("%v ", a*rn)
		rn *= r
		time.Sleep(5 * time.Millisecond)
	}
}

func main() {
	go printDeret(1, 2, 1, 4)
	go printChar("Halo")

	time.Sleep(50 * time.Millisecond)
}
