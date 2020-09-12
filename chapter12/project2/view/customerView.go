package main

import(
	"fmt"
	"go_code/chapter12/project2/service"
	"go_code/chapter12/project2/model"
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
//获取用户输入，构造新的客户信息，添加到管理系统中
func (cv *customerView)add()  {
	fmt.Println("-----------------------添加客户----------------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	gendar := ""
	fmt.Scanln(&gendar)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanln(&email)

	customer := model.NewCustomerWithoutId(name, gendar, age, phone, email)
	if cv.service.Add(customer){
		fmt.Println("-----------------------添加完毕----------------------")
	}else{
		fmt.Println("添加失败！")
	}
}
//删除客户
func (cv *customerView)deleta()  {
	fmt.Println("-----------------------删除客户----------------------")
	fmt.Print("请输入待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1{
		return
	}
	for {
		fmt.Print("确认是否删除[y/n]:")
		sure := ""
		fmt.Scanln(&sure)
		if sure == "n"{
			return 
		} else if sure == "y"{
			if cv.service.Delete(id){
				fmt.Println("-----------------------删除完毕----------------------")
			} else{
				fmt.Println("删除失败：客户ID不存在")
			}
			return	
		}
	}
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
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			cv.deleta()
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

