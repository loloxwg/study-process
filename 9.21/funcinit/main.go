package main

import (
	utils2 "daily-study/9.21/funcinit/utils"
	"fmt"
)

var age = test()

func test() int {
	fmt.Println("teat()")
	return 90
}

//为了看到全局变量是先被初始化的，我们这里先写函数
func init() {
	fmt.Println("init()")
}

//init函数,通常可以在init函数中完成初始化工作
func main() {
	fmt.Println("main()...age=", age)
	fmt.Println("Age=", utils2.Age, "Name=", utils2.Name)

}
