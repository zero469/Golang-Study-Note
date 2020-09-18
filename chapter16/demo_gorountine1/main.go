package main

import(
	"fmt"
	"sync"
	"time"
)


var(
	fac map[int]int = make(map[int]int)
	//全局读写锁
	lock sync.RWMutex
)


func test(n int)  {
	res := 1
	for i := 1; i <= n; i++{
		res *= i
	}
	lock.Lock()
	fac[n] = res
	lock.Unlock()
}

func main()  {
	for i := 1; i <= 20; i++{
		go test(i)
		
	}
	time.Sleep(time.Second)
	lock.RLock()
	for i, v := range(fac){
		fmt.Printf("%v的阶乘 = %v\n", i, v)
	}
	lock.RUnlock()
}