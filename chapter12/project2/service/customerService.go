package service

import (
	"go_code/chapter12/project2/model"
	"fmt"
)

//CustomerService 提供增删查改
type CustomerService struct {
	customers []model.Customer
	//当前客户数
	customerNum int
}

//List 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

//Add 增加新的customer到切片中
func (cs *CustomerService) Add(customer model.Customer) bool {
	//分配id
	cs.customerNum++
	customer.ID = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

//Update 修改指定id客户的信息
func (cs *CustomerService)Update(id int, name, gender, age, phone, email string) bool {
	index := cs.FindByID(id)
	
	if name != ""{
		cs.customers[index].Name = name
	}
	
	if gender != ""{
		cs.customers[index].Gender = gender
	}
	if age != ""{
		//从字符串中获取年龄信息
		var ageNum = 0
		fmt.Sscanln(age, &ageNum)
		if ageNum != 0{
			cs.customers[index].Age = ageNum	
		}
	}
	if phone != ""{
		cs.customers[index].Phone = phone
	}
	if email != ""{
		cs.customers[index].Email = email 
	}
	return true

}


//Delete 删除指定id的客户
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindByID(id)
	if index == -1 {
		return false
	}
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

//FindByID 根据id返回下标，如果没有该客户则返回-1
func (cs *CustomerService) FindByID(id int) int {
	for index, val := range cs.customers {
		if val.ID == id {
			return index
		}
	}
	return -1
}

//GetInfoByID 根据id获取客户信息
func (cs *CustomerService)GetInfoByID(id int) *model.Customer {
	for _, val := range cs.customers {
		if val.ID == id {
			return &val
		}
	}
	return nil
}

//NewCustomerService 创建新的CustomerService实例
func NewCustomerService() *CustomerService {
	cs := &CustomerService{}
	cs.customerNum = 1
	customer := model.NewCustomer(cs.customerNum, "张三", "男", 20, "123", "zhangsan@qq.com")
	cs.customers = append(cs.customers, customer)
	return cs
}