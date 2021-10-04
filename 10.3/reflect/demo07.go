package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	f1:=fun1
	value := reflect.ValueOf(f1)
	value2:=reflect.ValueOf(fun2)
	value3:=reflect.ValueOf(fun3)
	fmt.Printf("type: %s kind: %s\n",value.Type(), value.Kind())
	fmt.Printf("type: %s kind: %s\n",value2.Type(), value2.Kind())
	fmt.Printf("type: %s kind: %s\n",value3.Type(), value3.Kind())

	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(100),reflect.ValueOf("zxw")})

	call := value3.Call([]reflect.Value{reflect.ValueOf(100000), reflect.ValueOf("sb")})
	fmt.Println(call[0].Type(),call[0].Kind())

	s := call[0].Interface().(string)
	fmt.Printf("%T\n",s)

}


func fun1(){
	fmt.Println("fun1...无参数")
}

func fun2(i int,s string){
	fmt.Println("fun2...有参数",i,s)
}

func fun3(i int,s string) string {
    fmt.Println("fun3...有参数，有返回值")
	return s+strconv.Itoa(i)
}