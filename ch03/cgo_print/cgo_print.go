package main

// #include <stdio.h>
// #include <malloc.h>
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}

func main() {
	Print("Hello, from Go calling C!\n")
}
