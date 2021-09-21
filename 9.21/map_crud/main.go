package main

import "fmt"

func main() {

	cities := make(map[string]string)

	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"

	fmt.Println("last:", cities)

	cities["no1"] = "shanxi"
	//因为这个no1存在 后添加直接修改
	fmt.Println("now:", cities)

	//因为 no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	//演示删除
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	//如果希望一次性删除所有的key
	//1. 遍历所有的key,如何逐一删除 [遍历]
	//2. 直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)

	val, ok := cities["no1"]
	if ok {
		fmt.Println("有key no1", val)
	} else {
		fmt.Println("没有 key:no1 🐛")
	}
}
