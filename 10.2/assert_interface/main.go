package main

import (
	"fmt"
	"math"
)

func main() {
	var t1 Triangle = Triangle{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var c1 Circle = Circle{radius: 4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())
	fmt.Println(c1.radius)

	var s1 Shape
	s1 = t1
	fmt.Println(s1.peri())
	fmt.Println(s1.area())

	var s2 Shape
	s2 = c1
	fmt.Println(s2.peri())
	fmt.Println(s2.area())

	testShape(t1)
	testShape(c1)
	testShape(s1)

	getType(t1)
	getType(s1)

	var t2 *Triangle = &Triangle{
		3, 4, 5,
	}
	fmt.Printf("s:%T,%p,%p\n", t2, &t2,t2)
	getType(t2)
}

func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok {
		fmt.Println("是三角形，三边是", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("原型", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("ins:%T,%p,%p\n", ins, &ins,ins)
		fmt.Printf("s:%T,%p,%p\n", s, &s,s)
	}
}

func getType2(s Shape){
	switch ins:=s.(type) {
	case Triangle:
		fmt.Println("是三角形，三边是", ins.a, ins.b, ins.c)
	case Circle:
		fmt.Println("原型", ins.radius)
	case *Triangle:
		fmt.Println("三角形结构体指针")
	}

}
func testShape(s Shape) {
	fmt.Printf("%.2f %.2f", s.peri(), s.area())
}

type Shape interface {
	peri() float64
	area() float64
}

type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}
