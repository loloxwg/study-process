package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
}

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

//引用类型必须必须先分配内存
func main() {
	var p1 Person
	fmt.Println(p1)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil {
		fmt.Println("ok3")
	}
	p1.slice = make([]int, 10) //引用类型必须必须先分配内存
	p1.slice[0] = 100000

	p1.map1 = make(map[string]string) //引用类型必须必须先分配内存
	p1.map1["key1"] = "tom~"

	fmt.Println(p1)

	var monster1 Monster
	monster1.Name = "niumowang"
	monster1.Age = 100000

	monster2 := monster1 //结构体是值类型默认值拷贝
	monster2.Name = "shejing"

	fmt.Println(monster1)
	fmt.Println(monster2)

	//声明方式
	//方式一，直接
	var person Person
	fmt.Println(person)
	//方式二
	var person2 Person = Person{
		"marry",
		200,
		[5]float64{},
		nil,
		nil,
		nil,
	}
	fmt.Println(person2)
	//方式3  &
	var p3 *Person = new(Person)
	p3.Name = "zxw"   //都可以
	(*p3).Name = "zx" //都可以
	fmt.Println(*p3)
	//go 的设计者再底层会对 p3.Name = "zxw"处理
	//  p3加上取值运算 * 语法糖

	//方式4{}
	var person4 *Person = &Person{}
	person4.Name = "zzzzzz"
	(*person4).Name = "xxxxxxxx"
	fmt.Println(*person4)
	////  p3加上取值运算 * 语法糖

}
