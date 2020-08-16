package main

import (
	"fmt"
)

func fiboSli(n int) []uint64 {
	fibo := make([]uint64, n)
	fibo[0] = 1
	fibo[1] = 1
	for i := 2; i < n; i++ {
		fibo[i] = fibo[i-1] + fibo[i-2]
	}
	return fibo
}

func main() {

	str := "hello@guigu" //本质是一个切片，引用一个byte数组

	fmt.Printf("%p %T\n", &str, str)

	slice := str[6:]

	str01 := []byte(str)

	str01[0] = '+'

	str = string(str01)
	fmt.Println(str)
	fmt.Println(str01)
	fmt.Println(slice)

	fmt.Println(fiboSli(100))

}
