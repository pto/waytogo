package main

import "fmt"

func main() {
	var a [2600000]int
	a[len(a)-1] = 42
	var b [2600000]int
	b[len(b)-1] = 123
	var c [2600000]int
	c[len(c)-1] = 987
	fmt.Println(a[len(a)-1])
}
