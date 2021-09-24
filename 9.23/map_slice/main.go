package main

import "fmt"

func main() {
	var monster []map[string]string

	fmt.Printf("%T\n", monster)

	monster = make([]map[string]string, 3) //

	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "卡比"
		monster[0]["sex"] = "male"
		monster[0]["xxxx"] = "xxx"
	}

	if monster[2] == nil {
		monster[2] = make(map[string]string, 100)
		monster[2]["name"] = "卡比"
		monster[2]["sex"] = "male"
		monster[2]["xxxx"] = "xxx"
	}
	fmt.Printf("%v", monster)
}
