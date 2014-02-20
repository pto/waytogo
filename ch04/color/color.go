package main

import "fmt"

type Color int

const (
	Red Color = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

func main() {
	var c1 Color = Red
	// var c2 int = 0
	fmt.Println("c1 is", c1)
	fmt.Println("c1 == Red is", c1 == Red)
	fmt.Println("c1+1 is", c1+1)
	// ERROR: fmt.Println("c1 == c2 is", c1 == c2)
	// ERROR: fmt.Println("c1+c2 is", c1+c2)
}
