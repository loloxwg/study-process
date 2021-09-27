package main

import "fmt"

type Visitor struct {
	Name string
	Age  int
}

//String 方法重写
func (visitor *Visitor) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", visitor.Name, visitor.Age)
	return str
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("NO!")
		return
	}
	if visitor.Age >= 18 {
		fmt.Printf("visitor name= %v,age= %v price=20$", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("visitor name= %v,age= %v price=0$", visitor.Name, visitor.Age)
	}
}

func main() {
	var v Visitor
	v.Age = 18
	v.Name = "zxw"
	fmt.Println(v)
	fmt.Println(&v)
	for {
		fmt.Println("input your visitor name")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("exit")
			break
		}
		fmt.Println("input your age")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
