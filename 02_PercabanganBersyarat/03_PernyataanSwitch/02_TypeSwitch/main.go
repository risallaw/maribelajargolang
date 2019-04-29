package main

import "fmt"

func getType(i interface{}) string {
	switch i.(type) {
	case nil:
		return "nil"
	case int:
		return "integer"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println(getType("Hello"))
	fmt.Println(getType(13))
	fmt.Println(getType(2.85))
	fmt.Println(getType(false))
}
