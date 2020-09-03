package main

import(
	"fmt"
	"go_code/chapter11/test_encapsulate/model"
)

func main()  {
	account := model.NewAccount("123456", "111111", 100)
	fmt.Println(account)
	account.Query("111111")
}