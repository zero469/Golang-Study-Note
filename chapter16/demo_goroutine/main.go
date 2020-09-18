package main

import(
	"fmt"
	"time"
	"strconv"
	"runtime"
)

func test()  {
	for i := 0; i < 10; i++{
		fmt.Println("hello word " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main()  {

	go test() //开启一个协程	
	fmt.Println(runtime.NumCPU())
	for i := 0; i < 6; i++{
		fmt.Println("hello main " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}