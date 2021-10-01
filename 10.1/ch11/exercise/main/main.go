package main

import (
	"daily-study/10.1/ch11/exercise/model"
	"fmt"
)

func main() {
	account := model.NewAccount("account", "123456", 40)
	if account != nil {
		fmt.Println("successfully created account")
	} else {
		fmt.Println("failed err")
	}

	account.Query("123456")
	account.Deposite(20, "123456")
	account.Query("123456")
	account.SetAccount("xxxxx")
	account.GetAccount()

}
