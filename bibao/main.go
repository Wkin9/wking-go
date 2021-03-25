package main

import "fmt"

func main() {

	ret := add()(200)

	fmt.Println(ret)

	f1(f3(f2, 10, 20))

}

func add() func(int) int {
	x := 100
	return func(y int) int {
		return x + y
	}
}

func f1(f func()) {

	fmt.Println("f1函数执行。。。")
	f()

}

func f2(x, y int) {

	fmt.Printf("f2函数执行。。。x+y=%d \n", x+y)

}

func f3(f func(int, int), x, y int) func() {

	fmt.Println("f3函数执行了。。。。")
	ret := func() {
		f(x, y)
	}
	return ret
}
