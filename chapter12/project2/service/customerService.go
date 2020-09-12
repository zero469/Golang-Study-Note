package service
import(
	"go_code/chapter12/project2/model"
)
//CustomerService 提供增删查改
type CustomerService struct{
	customers []model.Customer
	//当前客户数
	customerNum int

}
//List 返回客户切片
func (cs *CustomerService) List() [] model.Customer{
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
//Delete 删除指定id的客户
func (cs *CustomerService)Delete(id int)  bool{
	index := cs.findByID(id)
	if index == -1{
		return false
	}
	cs.customers = append(cs.customers[:index], cs.customers[index + 1:]...)
	return true
}
//根据id返回下标，如果没有该客户则返回-1
func (cs *CustomerService)findByID(id int) int{
	for index, val := range(cs.customers){
		if val.ID == id{
			return index
		}
	}
	return -1
}
//NewCustomerService 创建新的CustomerService实例
func NewCustomerService() *CustomerService{
	cs := &CustomerService{}
	cs.customerNum = 1
	customer := model.NewCustomer(cs.customerNum, "张三", "男", 20, "123", "zhangsan@qq.com")
	cs.customers = append(cs.customers, customer)
	return cs
}