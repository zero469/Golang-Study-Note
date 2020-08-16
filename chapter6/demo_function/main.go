package main

import(
	"fmt"
)

// func cal(n1 float64, n2 float64, op byte) float64{
// 	switch op {
// 	case '+':
// 		return n1 + n2
// 	case '-':
// 		return n1 - n2
// 	case '*':
// 		return n1 * n2
// 	case '/':
// 		return n1 / n2
// 	default:
// 		fmt.Println("error : worong operator")
// 		return 0
// 	}
// }

func fibo(n int) int {
	if (n == 1 || n == 2){
		return 1
	}
	return fibo(n - 1) + fibo(n - 2)
}
//基本数据类型和数组是值拷贝
func test(n1 int)  {
	n1 += 10
	fmt.Println(n1)
	
}

func test1(n1 *int){
	*n1 += 10
	fmt.Println(*n1)
}

func reverse(arr [4]int, len int)  {
	for i := 0; i < len / 2; i++{
		temp := arr[i]
		arr[i] = arr[len - i - 1]
		arr[len - i - 1] = temp
	}
	fmt.Println(arr)
}
//自定义类型
type opfunc func(int, int) int

func sum(num1 int, num2 int) int{
	return num1 + num2
}
//返回值命名
func cal(op opfunc, num1 int, num2 int) (ret int) {
	ret = op(num1, num2)
	return 
}

//可变参数，只能放在最后面
func sumArgs(num int, nums... int) int {
	sum := num
	for _, val := range nums{
		sum += val
	}
	return sum
}

func main()  {
	// fmt.Println(cal(1.2, 1.3, '+'))
	// fmt.Println(cal(1.2, 1.3, '-'))
	// fmt.Println(cal(1.2, 1.3, '*'))
	// fmt.Println(cal(1.2, 1.3, '/'))
	// fmt.Println(cal(1.2, 1.3, '='))
	// for i := 1; i < 8; i++{
	// 	fmt.Println(fibo(i))
	// }

	// var n1 = 10
	// test(n1)
	// fmt.Println(n1)
	// test1(&n1)
	// fmt.Println(n1)
	// var arr = [4]int{1, 2, 3, 4}
	// reverse(arr, 4)//参数传递
	// fmt.Println(arr)

// 	fmt.Println(cal(sum, 1, 2))
// 	fmt.Println(sumArgs(1, 2, 3, 4, 5))
		
}