package main

import "fmt"

func main() {
	//第一种使用方式
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "宋江" //ok?
	a["no2"] = "吴用" //ok?
	a["no1"] = "武松" //ok?
	a["no3"] = "吴用" //ok?
	fmt.Println(a)

	//第二种方式  声明直接赋值
	cities := make(map[string]string)

	cities["no1"] = "beijing"
	cities["no2"] = "tianjin"
	cities["no3"] = "shanghai"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "songjiang",
		"hero2": "songjiba",
	}
	heroes["hero4"] = "zxwan"
	fmt.Println("heroes", heroes)

}
