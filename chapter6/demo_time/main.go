package main

import(
	"time"
	"fmt"
	"strconv"
)

func test03()  {
	str := ""
	for i := 1; i < 100000; i++{
		str += strconv.Itoa(i)
	}
}

func main()  {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	
	
	//日期格式化
	fmt.Printf(now.Format("2006-01-02\n"))
	fmt.Printf(now.Format("2006 01 02\n"))


	//时间常量
	time.Sleep(2 * time.Microsecond)
	fmt.Println("ok")


	//
	fmt.Println(now.Unix())//返回秒时间戳
	fmt.Println(now.UnixNano())


	//函数运行时间
	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Println(end - start)
}