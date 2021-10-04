package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64=123456789
	fmt.Println(num)
	pointer:=reflect.ValueOf(&num)

	newValue:= pointer.Elem()

	fmt.Println(newValue.Type())
	fmt.Println(newValue.CanSet())

	newValue.SetFloat(2222.2)
	fmt.Println(num)


}
