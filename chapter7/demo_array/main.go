package main

import(
	"fmt"
	"unsafe"
	"math/rand"
	"time"
)

func test01(arr [3]int)  {
	fmt.Printf("%p\n",&arr)
	arr[0]++
	arr[1]++
	arr[2]++
}

func test02(arr *[3]int)  {
	fmt.Printf("%p\n",arr)
	(*arr)[0]++
	(*arr)[1]++
	(*arr)[2]++
	
}

func main()  {
	var intArr [3]int
	fmt.Println(intArr)
	fmt.Printf("%p\n",  &intArr)
	fmt.Println(unsafe.Sizeof(intArr[0]))
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// var score [5]float64
	// for i := 0; i < len(score); i++{
	// 	fmt.Scanln(&score[i])
	// }
	// fmt.Println(score)

	//初始化
	var arr01 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr01)
	var arr02 = [3]int{1, 2, 3}
	fmt.Println(arr02)
	var arr03 = [...]int{1, 2, 3}
	fmt.Println(arr03)
	var arr04 = [...]int{0 : 200, 20: 1}//可省略前面的
	fmt.Println(arr04)

	//for range遍历
	for i, val := range arr04{
		fmt.Println(val)
		fmt.Println(i,":",arr04[i])
	}


	//数组值拷贝
	var testArr [3]int
	fmt.Println(testArr)
	fmt.Printf("%p\n", &testArr)
	test01(testArr)
	fmt.Println(testArr)

	//传引用
	test02(&testArr)
	fmt.Println(testArr)

	//练习
	rand.Seed(time.Now().UnixNano())
	var exe01 [5]int

	for i := 0; i < len(exe01); i++{
		exe01[i] = rand.Intn(100)//生成0-100的随机整数
	}
	fmt.Println(exe01)
	for i := len(exe01) - 1; i >= 0; i--{
		fmt.Printf("%v ",exe01[i])
	}
	for i := 0; i < len(exe01); i++{
		defer fmt.Printf("%v ",exe01[i])//defer到最后反转打印
	}
	fmt.Println()
}