package main

import (
	"fmt"
	"strconv"
)


type Human struct {
	name string
	age int
	phone string
}

func (h Human) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
	Bob:=Human{
		name: "Bob",
		age: 39,
		phone: "123456789",
	}
	fmt.Println(Bob)
}


