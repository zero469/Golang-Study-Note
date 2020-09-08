package model

import(
	"fmt"
)


type Customer struct{
	ID int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

func (c Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",
			c.ID, c.Name, c.Gender, c.Age, c.Phone, c.Email)
	return info
}

func NewCustomer(id int, name string, gendar string,
	age int, phone string, email string)  Customer{
	return Customer{
		ID : id,
		Name : name,
		Gender : gendar,
		Age : age,
		Phone : phone,
		Email : email,
	}
}