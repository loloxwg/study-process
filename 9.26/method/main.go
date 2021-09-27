package main

import "fmt"

type Person1 struct {
	Name string
}

func (p Person1) test() {
	fmt.Println("test() name=", p.Name)
}
func (p Person1) speak() {
	fmt.Println(p.Name, "是个好人")
}
func (p Person1) jisuan() {
	res := 0
	for i := 0; i < 100000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是=", res)
}

func (p Person1) jisuan2(n int) {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算结果是=", res)
}

func (p Person1) getSum(a, b int) int {
	return a + b
}

func main() {
	var p Person1
	p.Name = "tom"
	p.test()
	p.speak()
	p.jisuan()
	p.jisuan2(100000)
	res := p.getSum(1, 2)
	fmt.Println(res)
}
