package main

import(
	"fmt"
	"time"
)

func test(){
	var mp map[int]string
	//在协程中使用recover捕获panic
	//防止因为一个协程的崩溃导致整个进程崩溃
	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println("test函数 error：",err)
		}
	}()
	mp[1] = "a"
}

func main()  {
	go test()
	time.Sleep(time.Second)
}