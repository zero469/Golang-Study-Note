package main

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str := "北京123"
	//返回字节数，汉字三个字节，字母一个字节
	fmt.Println(len(str))

	//遍历含中文的字符串
	r := []rune(str)
	for _,val := range r{
		fmt.Println(string(val))
	}


	//字符串转整数
	n, err := strconv.Atoi("111111111111111111")
	if err != nil{
		fmt.Println("error: ",err)
	}else{
		fmt.Println(n)
	}

	//整数转字符串
	str = strconv.Itoa(1233)
	fmt.Println(str)

	//进制转换，返回字符串
	fmt.Println(strconv.FormatInt(100, 16))

	//不区分大小写比较
	fmt.Println(strings.EqualFold("abc", "ABC"))

		
}