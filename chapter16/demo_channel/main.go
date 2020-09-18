package main

import(
	"fmt"
)

func main()  {
	//管道示例
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("%v, %T\n", intChan, intChan)

	//写管道
	intChan<- 10
	num := 100
	intChan<- num
	intChan<- 10000
	//超出容量会死锁


	//读管道
	var num2 int
	num2 = <-intChan
	fmt.Println(num2)

	fmt.Println(len(intChan), cap(intChan))
}