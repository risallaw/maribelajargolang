package main

import (
	"fmt"
	"time"
)

func halo() {
	fmt.Println("Halo.")
}

func main() {
	fmt.Println("Fungsi main dimulai.")
	go halo()

	time.Sleep(10 * time.Millisecond)
	fmt.Println("Fungsi main selesai.")
}
