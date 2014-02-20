// hello_world2 is a demonstration program that prints a greeting and calls demo.Thing().
package main

import (
	"wayofgo/ch04/hello_world2/demo"
	"fmt" // Package implementing formatted I/O.
)

var _, _ = fmt.Println("var count in package main")

func main() {
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	demo.Thing()
	fmt.Println(part2)
}

func init() {
	fmt.Println("main.init()")
}
