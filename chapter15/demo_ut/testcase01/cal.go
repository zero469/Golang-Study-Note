package main

func AddUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++{
		res += i
	}
	return res
}

func Sub(a ,b int) int {
	res := a - b
	if res <= 0{
		return 0
	}
	return res
}

func main(){
	AddUpper(10)
}

