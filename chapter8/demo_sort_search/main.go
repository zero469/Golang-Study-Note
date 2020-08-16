package main

import(
	"fmt"
	"time"
)

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++{
		flag := false
		for j := len(nums) - 1; j > i; j--{
			if(nums[j] < nums[j - 1]){
				temp := nums[j - 1]
				nums[j - 1] =nums[j]
				nums[j] = temp
				flag = true
			}
		}
		if(!flag){
			return
		}
	}
}

func main()  {
	var nums []int =[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	start := time.Now().UnixNano()
	bubbleSort(nums)
	end := time.Now().UnixNano()
	fmt.Println(nums, "time is :", end - start)
}