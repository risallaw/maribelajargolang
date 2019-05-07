package main

import "fmt"

func main() {
	c := make(chan int)
	fmt.Printf("Tipe dari c adalah %T\n", c)
	fmt.Printf("Nilai dari c adalah %v\n", c)
}
