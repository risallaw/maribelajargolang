package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, err1 := strconv.ParseBool("true")
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(b)

	f, err2 := strconv.ParseFloat("3.1415", 64)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(f)

	i, err3 := strconv.ParseInt("2019", 10, 0)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(i)
}
