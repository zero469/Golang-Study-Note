package main

import(
	"fmt"
)

func main()  {
	var mp map[string]string
	mp = make(map[string]string, 10)
	mp["no1"] = "no1"
	mp["no2"] = "no2"
	mp["no3"] = "no3"
	mp["no4"] = "no4"
	fmt.Println(mp)

	mp01 := make(map[int]int)
	fmt.Println(len(mp01))
	for i := 0; i < 10000; i++{
		mp01[i] = i
	}
	fmt.Println(mp01)

	mp02 := map[int]int{
		1 : 1,
		2 : 2,
	}
	fmt.Println(mp02)
}