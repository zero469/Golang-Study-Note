package main

import(
	"runtime"
	"fmt"
)

func main()  {
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1)
}