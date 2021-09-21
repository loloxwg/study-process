package main

import (
	"fmt"
	"strings"
)

// AddUpper 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello world"
	return func(x int) int {
		n = n + x
		str = str + string(36) //=>36='$'
		fmt.Println(str)
		return n
	}
}

// 1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) == false {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	//测试 makeSuffix 的使用
	//返回一个闭包
	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后：", f2("world"))
	fmt.Println("文件名处理后：", f2("a.jpg")) //
	fmt.Println("文件名处理后：", f2("c.j"))

}
