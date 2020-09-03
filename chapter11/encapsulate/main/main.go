package main

import(
	"fmt"
	"go_code/chapter11/encapsulate/model"
)

func main()  {
	p := model.NewPerson("smith")
	p.SetAge(18)
	p.SetSalary(5000)
	fmt.Println(p, "age = ",p.GetAge(), "salary = ", p.GetSalary())
}