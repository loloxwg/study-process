package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

//`json:"name"`就是struct_tag
func main() {
	monster1 := Monster{
		"niumowang",
		500,
		"bajiaoshan",
	}
	jsonStr, err := json.Marshal(monster1)
	if err != nil {
		fmt.Println("json.Marshal Error", err)
	}
	fmt.Println("jsonStr", string(jsonStr))

}
