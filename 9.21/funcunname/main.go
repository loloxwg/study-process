package main

//匿名函数
import "fmt"

var (
	fun1 = func(n1, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//1.直接匿名函数，使用一次
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(10, 20)

	fmt.Println("res1:", res1)
	//2.函数变量匿名函数，多次
	a := func(n1, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 30)
	res3 := a(1, 40)
	fmt.Println("res2:", res2)
	fmt.Println("res3:", res3)

	//全局匿名函数,不常用
	res4 := fun1(1, 20)
	fmt.Println("res4:", res4)
}
