package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start.")

	n := make(chan string, 5)
	n <- "hello"

	fmt.Println("End.")
}
