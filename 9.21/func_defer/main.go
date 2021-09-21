package main

import (
	"fmt"
)

func sum(n1, n2 int) int {

	//当执行到defer时，会将defer后面的语句和  相关参数  压入到独立的栈中(defer栈)
	//当函数执行完毕以后，在从defer栈，按照先入后出的方式出栈
	defer fmt.Println("ok1 n1=", n1) // 3.   1
	defer fmt.Println("ok2 n2=", n2) // 2.   2

	//加一句话
	n1++
	n2++

	res := n1 + n2
	fmt.Println("ok3 res=", res) // 1.    3
	return res

}

func main() {
	res := sum(1, 2)
	fmt.Println("res=", res) // 4.   3====>5
}
