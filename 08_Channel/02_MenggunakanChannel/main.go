package main

import "fmt"

func luasPersegi(s chan int) {
	sisi := <-s
	fmt.Printf("Luas persegi %v persegi.\n", sisi*sisi)
}

func main() {
	fmt.Println("Start.")

	s := make(chan int)
	go luasPersegi(s)

	s <- 4
	fmt.Println("End.")
}
