// Package demo is a demonstration package to test out godoc.
//
// Pretty much all this package does is provide a place to put some documentation comments,
// so we can see how they work.
package demo

import "fmt"

// count is the number of characters in the Println, and err should be nil.
var count, err = fmt.Println("var in package demo")
var _,_ = fmt.Println("unnamed var in package demo")

// Thing is a function that prints "Thing".
//
// It does this incredible feat by calling fmt.Println with a string. Remarkably,
// the contents of that string is "Thing". Thus, the resulting output is generated.
func Thing() {
	fmt.Println("Thing")
}

func init() {
	fmt.Println("init in package demo")
}
