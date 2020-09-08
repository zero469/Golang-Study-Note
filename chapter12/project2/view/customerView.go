package main

import(
	"fmt"
	"go_code/chapter12/project2/service"
)
type customerView struct{
	//操作码
	opCode string
	//是否循环显示主菜单
	loop bool
	service *service.CustomerService
}
//获取customer信息
func (cv *customerView)list()  {
	customers := cv.service.List()
	//显示信息
	fmt.Println("-----------------------客户信息----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, val := range(customers){
		fmt.Println(val.GetInfo())
	}
	fmt.Println("-----------------------显示结束----------------------")
	
}

//显示主菜单
func (cv *customerView) mainMenu()  {
	for{
		fmt.Println("-------------------客户信息管理软件-------------------")
		fmt.Println("                   1 添 加 客 户")
		fmt.Println("                   2 修 改 客 户")
		fmt.Println("                   3 删 除 客 户")
		fmt.Println("                   4 客 户 列 表")
		fmt.Println("                   5 退       出")

		fmt.Print("请选择(1-5):")
		fmt.Scanln(&cv.opCode)
		switch cv.opCode{
		case "1":
			fmt.Println("添加客户")
		case "2":
			fmt.Println("修改客户")
		case "3":
			fmt.Println("删除客户")
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Println("输入有误")
		}
		if !cv.loop{
			break;
		}
	}
	fmt.Println("Bye bye!")
}

func main()  {
	
	newView := &customerView{
		opCode : "",
		loop : true,
	}
	newView.service = service.NewCustomerService()
	newView.mainMenu()
}

