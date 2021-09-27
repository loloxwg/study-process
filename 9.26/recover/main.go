package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生了错误")
		}
	}()
	var myMap map[int]string
	///error 显然这里没有为map分配内存
	myMap[0] = "golang"
}
func main() {
	go sayHello()

	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok", i)
		time.Sleep(time.Second)
	}
}
