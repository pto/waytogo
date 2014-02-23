package main

import "fmt"

func main() {
	i := 5
	switch i {
	case 5:
		fmt.Println(">4")
		fallthrough
	case 4:
		fmt.Println(">3")
		fallthrough
	case 3:
		fmt.Println(">2")
		fallthrough
	case 2:
		fmt.Println(">1")
		fallthrough
	case 1:
		fmt.Println(">0")
		fallthrough
	default:
		fmt.Println("in default")
	}
}
