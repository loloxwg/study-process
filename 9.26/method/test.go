package main

import "fmt"

type Person struct {
	Name string
}

func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03()=", p.Name)
}
func (p *Person) test04() {
	p.Name = "marry"
	fmt.Println("test03()=", p.Name)
}
func main() {
	p := Person{
		"tom",
	}
	test01(p)
	test02(&p)
	p.test04()
	p.test03()
}
