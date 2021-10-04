package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64=123456789
	value:=reflect.ValueOf(num)

	convertValue:=value.Interface().(float64)
	fmt.Println(convertValue)

	pointer:=reflect.ValueOf(&num)
	convertPinter:=pointer.Interface().(*float64)
	fmt.Println(convertPinter)
}
