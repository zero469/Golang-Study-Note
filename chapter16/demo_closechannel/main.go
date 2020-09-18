package main

import(
	"fmt"
)

func main()  {
	intChan := make(chan int, 100)
	for i := 0; i < 100; i++{
		intChan<- i * 2
	}

	close(intChan)
	//遍历时，如果没有close管道，会出现死锁，但是可以把数据取完
	//如果close管道就可以正确读取
	for v := range intChan{
		fmt.Println(v)
	}
}