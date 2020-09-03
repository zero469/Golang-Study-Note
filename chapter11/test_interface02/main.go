package main

import(
	"fmt"
	"sort"
	"math/rand"
)
//Hero 自定义类型
type Hero struct{
	Name string
	Age int
}
//HeroSlice 定义类型的切片
type HeroSlice []Hero

//实现sort.Sort所需的Interface
func (hs HeroSlice)Len() int {
	return len(hs)
}
func (hs HeroSlice)Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice)Swap(i, j int){
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main()  {
	var intSlice = []int{5 ,4, 3, 2, 1}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heros HeroSlice
	for i := 0; i < 10; i++{
		hero := Hero{
			Name : fmt.Sprintf("英雄：%d", rand.Intn(100)),
			Age : rand.Intn(100),
		}
		heros = append(heros, hero)
	}
	fmt.Println(heros)
	sort.Sort(heros)
	fmt.Println(heros)
}