package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start.")

	n := make(chan string)
	n <- "hello"

	fmt.Println("End.")
}
