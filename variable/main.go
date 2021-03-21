package main

import "fmt"

func main() {

	fmt.Println("hello world")
	fmt.Println("--------------------------------")
	variable()
	fmt.Println()
	variableInit()
}

func variable() {

	var i int
	var s string
	fmt.Printf("i=%d s=%q", i, s)
}
func variableInit() {

	var i, j int = 1, 2
	var s string = "hello"
	fmt.Printf("i=%d j=%d s=%q", i, j, s)
}
