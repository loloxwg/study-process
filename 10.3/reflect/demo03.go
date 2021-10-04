package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
	Sex string
}

func (p Person) Say(msg string)  {
	fmt.Println("hello",msg)
}

func (p Person) Println()  {
	fmt.Println("name",p.Name,"age",p.Age,"sex",p.Sex)
}

func main() {
	p1:=Person{
		"zxw",
		30,
		"male",
	}
	GetMessage(p1)

}

func GetMessage(input interface{}){
	getType := reflect.TypeOf(input)
	fmt.Println("getType", getType.Name())//Person
	fmt.Println("get Kind",getType.Kind())//struct

	getValue:=reflect.ValueOf(input)
	fmt.Println("get all fields",getValue)

	for i := 0; i < getType.NumField(); i++ {
		field:=getType.Field(i)
		value:=getValue.Field(i).Interface()
		fmt.Printf("字段名称 %s 字段类型 %s 字段数值 %v\n",field.Name,field.Type,value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		method:=getType.Method(i)
		fmt.Printf("方法名字 %s 方法类型 %s\n",method.Name,method.Type)
	}
}