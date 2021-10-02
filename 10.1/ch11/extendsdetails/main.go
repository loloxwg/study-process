package main

import "fmt"

type A struct {
	Name string
	age int
}

func (a A)SayOK()  {
	fmt.Println("A SayOK",a.Name)
}

func (a A) hello() {
	fmt.Println("A Hello",a.Name)
}

type B struct {
	A
}

func main() {
	var b B
	b.A.Name="zxw"
	b.A.age=123456789
	b.A.SayOK()
	b.A.hello()
	b.SayOK()
	m := map[byte]int{}
	fmt.Println(m)
}



