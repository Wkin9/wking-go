package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {

	fmt.Println("invar")
	complx()
	triangle(3, 4)
	triangle(5, 12)
}

func complx() {

	a := 3 + 4i // 3^2+4^2=5^2
	fmt.Println(cmplx.Abs(a))

}

func triangle(a, b int) {

	c := math.Sqrt(float64(a*a + b*b))
	fmt.Printf("直角边长为:%d ,直角边长为:%d, 三角形的斜边是：%.0f", a, b, c)
	fmt.Println()

}
