// function_calls_function expected output is:
//
// GOG (no newline)
package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
