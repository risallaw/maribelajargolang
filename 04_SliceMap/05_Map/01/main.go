package main

import (
	"fmt"
	"strconv"
)

func main() {
	angka := map[int]string{
		0: "nol", 1: "satu",
		2: "dua", 3: "tiga",
		4: "empat", 5: "lima",
		6: "enam", 7: "tujuh",
		8: "delapan", 9: "sembilan",
	}

	var nStr string
	fmt.Print("Masukan deret angka: ")
	fmt.Scanln(&nStr)

	for _, el := range nStr {
		num, _ := strconv.Atoi(string(el))
		fmt.Printf("%s ", angka[num])
	}
	fmt.Println()
}
