package main

import "fmt"

func tambahKurang(a, b int) (tambah, kurang int) {
	tambah = a + b
	kurang = a - b
	return
}

func main() {
	var a, b int
	fmt.Print("Masukan nilai a dan b: ")
	fmt.Scanf("%d %d\n", &a, &b)

	t, k := tambahKurang(a, b)
	fmt.Printf("Tambah: %v\nKurang: %v\n", t, k)
}
