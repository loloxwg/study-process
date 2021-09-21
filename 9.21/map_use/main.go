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

	//案例
	/*
		课堂练习：演示一个key-value 的value是map的案例
		比如：我们要存放3个学生信息, 每个学生有 name和sex 信息
		思路:   map[string]map[string]string

	*/

	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "male"
	studentMap["stu01"]["address"] = "beijing"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "marry"
	studentMap["stu02"]["sex"] = "male"
	studentMap["stu02"]["address"] = "shanghai"

	studentMap["stu03"] = make(map[string]string, 3)
	studentMap["stu03"]["name"] = "zom"
	studentMap["stu03"]["sex"] = "male"
	studentMap["stu03"]["address"] = "shenzhen"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu03"])
	fmt.Println(studentMap["stu01"]["name"])

}
