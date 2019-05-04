package main

import "fmt"

type alamat struct {
	jalan string
	kota  string
}

type pegawai struct {
	nama   string
	alamat alamat
}

func main() {
	var p pegawai

	fmt.Print("Nama: ")
	fmt.Scanln(&p.nama)

	fmt.Print("Jalan: ")
	fmt.Scanln(&p.alamat.jalan)

	fmt.Print("Kota: ")
	fmt.Scanln(&p.alamat.kota)

	fmt.Println()
	fmt.Println(p)
}
