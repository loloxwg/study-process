package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := func(a,b int) int{
		ans:=a+b
		return ans
	}(1,222)

	fun:= func1()
	res1:=fun(1,222)

	fmt.Println(res,res1)


}
//1`

type myfun func(int,int)string


func func1()myfun{
	fun:=func(a,b int)string{
		s:=strconv.Itoa(a)+strconv.Itoa(b)
		return s
	}
	return fun
}