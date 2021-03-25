package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("const")
	func() {

	}()
	constFun()

}
func constFun() {
	const (
		filename = "我的世界"
		i, j     = 5, 12
	)
	var k int
	k = int(math.Sqrt(float64(i*i + j*j)))

	fmt.Println(filename, i, j)

	fmt.Println(k)

	fmt.Print(k, "\n")

	fmt.Printf("i:%d j:%d\n", i, j)

}
