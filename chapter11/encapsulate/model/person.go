package model

import(
	"fmt"
)
//非导出，不可被其他包引用
type person struct{
	name string
	age int
	salary float64
}

//提供工厂模式访问接口
func NewPerson(name string) *person {
	return &person{
		name : name,
	}
}


//提供set和get方法
func (p *person)SetAge(age int)  {
	if age > 0 && age < 150{
		p.age  = age
	} else {
		fmt.Println("年龄范围错误")
	}
}
func (p *person)GetAge() int {
	return p.age
}

func (p *person)SetSalary(salary float64)  {
	if salary >= 3000 && salary < 30000 {
		p.salary = salary
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person)GetSalary() float64 {
	return p.salary
}


