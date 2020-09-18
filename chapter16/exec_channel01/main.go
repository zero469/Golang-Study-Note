package main

import(
	"fmt"
)

func initData(numChan chan int)  {
	for i := 1; i <= 20000; i++{
		numChan<- i
	}
	close(numChan)
	fmt.Println("close numChan")
}

func getSum(numChan chan int, resNum chan int64)  {
	var res int64 = 0
	for {
		v, ok := <-numChan
		if !ok && v == 0{
			// fmt.Println("get sum done")
			break
		}
		res += int64(v)
	}
	resNum<- res
	fmt.Println("write res" , res)
}

func main()  {
	numChan := make(chan int)
	resChan := make(chan int64)
	// exitChan := make(chan bool)
	go initData(numChan)
	for i := 0; i < 8; i++{
		go getSum(numChan, resChan)
	}
	//需要确定resChan什么时候关闭
	for i := 0; i < 8; i++{
		v, ok := <-resChan
		if !ok{
			break
		}
		fmt.Printf("res[%v] = %v\n",i , v)
	}
	close(resChan)

}