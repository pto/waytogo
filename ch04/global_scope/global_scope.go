// global_scope expected output is:
//
// GOO (no newline)
package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}
