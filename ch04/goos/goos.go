package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// var goos string = os.Getenv("GOOS")
	var goos = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Println("runtime.Compiler is", runtime.Compiler)
	fmt.Println("runtime.GOARCH is", runtime.GOARCH)
	fmt.Println("runtime.NumCPU() is", runtime.NumCPU())
	fmt.Println("runtime.NumGgoCall() is", runtime.NumCgoCall())
	fmt.Println("runtime.NumGoroutine() is", runtime.NumGoroutine())
	fmt.Println("runtime.Version() is", runtime.Version())
}
