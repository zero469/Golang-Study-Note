package main

import(
	"fmt"
	"errors"
)
//捕获错误
func test()  {
	defer func ()  {
		err := recover()
		if err != nil{
			fmt.Println("err = ",err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}
//自定义错误
func readConf(name string) (err error)  {
	if name == "config.ini"{
		return nil
	} 
	return errors.New("读取文件错误")
}
func test02()  {
	err := readConf("config.ini")
	if err != nil{
		panic(err)
	}
	fmt.Println("test02 ok!")
}

func main()  {
	// test()
	test02()
	fmt.Println("ok")
}