package main

import "fmt"

func halo() {
	fmt.Println("Halo.")
}

func main() {
	fmt.Println("Fungsi main dimulai.")
	go halo()
	fmt.Println("Fungsi main selesai.")
}
