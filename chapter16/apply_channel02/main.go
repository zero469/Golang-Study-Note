package main

import(
	"fmt"
	"math"
	"time"
)

// func checkPrime(intChan, primeChan chan int, exitChan chan bool) {
// 	for{
// 		num, ok := <-intChan
// 		if !ok{
// 			//intChan的数据取完了，该协程退出
// 			break
// 		}
// 		sqrtNum := int(math.Sqrt(float64(num)))
// 		flag := true
// 		for i := 2; i <= sqrtNum; i++{
// 			if num % i == 0{
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag{
// 			primeChan<- num
// 		} 
// 	}
// 	exitChan<- true
// 	// fmt.Println("协程退出")
// }

func main(){
	// intChan := make(chan int, 1000)
	// primeChan := make(chan int, 2000)
	// exitChan := make(chan bool, 4)

	// start := time.Now().UnixNano()
	// go func(intChan chan int){
	// 	for i := 1; i < 8000; i++{
	// 		intChan<- i
	// 	}
	// 	close(intChan)
	// }(intChan)
	
	// //time.Sleep(time.Microsecond)
	// for i := 0; i < 8; i++{
	// 	go checkPrime(intChan, primeChan, exitChan)
	// }

	// //如果primeChan容量不够，则其会阻塞，因此exitChan为空，那么此时会产生死锁
	// //即消费primeChan的数据需要先给exitChan生产数据，但是往exitChan中生产数据需要等待
	// //primeChan生产完，如果primeChan容量不够，就会产生死锁
	// for i := 0; i < 8; i++{
	// 	<-exitChan
	// }
	// close(primeChan)
	// end := time.Now().UnixNano()
	// fmt.Println(end - start)
	// for prime := range primeChan{
	// 	fmt.Println(prime)
	// }
	
	//---------------------单线程对照组----------------------
	//			8协程    单线程
	//计算时间   1930400  3055300 //单位ns
	start := time.Now().UnixNano()
	primeMp := make(map[int]int)

	primeNum := 1
	for i := 1; i <= 8000; i++{
		flag := true;
		sqrtI := int(math.Sqrt(float64(i)))
		for j := 2; j <= sqrtI; j++{
			if i % j == 0{
				flag = false
				break
			}
		}
		if flag{
			primeMp[primeNum] = i
			primeNum++
		}
	}
	end := time.Now().UnixNano()
	fmt.Println(end - start)
	
	
}