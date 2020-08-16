package main

import(
	"fmt"
)
func main()  {
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}


	slice := intArr[1: 3]//引用数组的[1:3]

	fmt.Println(slice)
	fmt.Println("元素个数：", len(slice))
	fmt.Println("容量个数：", cap(slice))
	fmt.Printf("原地址：%p， 切片地址：%p\n",&intArr, &slice)
	fmt.Printf("数组第2个元素地址：%p， 切片第1个元素地址：%p\n",&intArr[1], &slice[0])
	slice[0] = 100//修改了slice也会同时修改其引用的数组
	fmt.Println(intArr, slice)



	//切片的使用
	fmt.Println("\n---------------------使用切片---------------------")
	var slice01 []int = make([]int, 4)
	fmt.Println(cap(slice01))
	fmt.Println(len(slice01))
	fmt.Println(slice01)

	var slice02 []int = []int{1, 2, 3}
	fmt.Println(cap(slice02))
	fmt.Println(len(slice02))
	fmt.Println(slice02)

	slice02 = append(slice02, 4, 5)
	fmt.Println(cap(slice02))
	fmt.Println(len(slice02))
	fmt.Println(slice02)

	slice03 := append(slice02, 6, 7)
	fmt.Println(cap(slice02))
	fmt.Println(cap(slice03))

	
}