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
func (cs *CustomerService) List() [] model.Customer{
	return cs.customers
}

func NewCustomerService() *CustomerService{
	cs := &CustomerService{}
	cs.customerNum = 1
	customer := model.NewCustomer(cs.customerNum, "张三", "男", 20, "123", "zhangsan@qq.com")
	cs.customers = append(cs.customers, customer)
	return cs
}