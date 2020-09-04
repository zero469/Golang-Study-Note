package utils

import(
	"fmt"
)

type familyAccount struct{
	//账户余额
	balance float64
	//收支金额
	money float64
	//收支明细
	note string
	//收支明细
	detial string
	//是否有收支行为
	hasUsed bool
	//是否退出
	quitFlag bool

}
//NewFamilyAccount 创建接口
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		balance : 0.0,
		money: 0.0,
		note : "",
		detial : "收支\t账户余额\t收支金额\t说  明\n",
		hasUsed : false,
		quitFlag : false,
	}
}

func (fa *familyAccount)showDetial()  {
	fmt.Println("--------------当前收支明细记录-----------------")
	if !fa.hasUsed{
		fmt.Println("当前还没收支明细....记一笔吧！")
	} else{
		fmt.Println(fa.detial)
	}
	fmt.Println("")
}

func (fa *familyAccount)income() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&fa.money)
	fmt.Print("本次收入说明：")
	fmt.Scanln(&fa.note)
	fa.hasUsed = true
	fa.balance += fa.money
	fa.detial += fmt.Sprintf("收入\t%8.2f\t%8.2f\t%v\n", fa.balance, fa.money, fa.note)
}
func (fa *familyAccount)pay() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance{
		fmt.Println("支出金额超过余额！")
		return
	}
	fmt.Print("本次支出说明：")
	fmt.Scanln(&fa.note)
	fa.hasUsed = true
	fa.balance -= fa.money
	fa.detial += fmt.Sprintf("支出\t%8.2f\t%8.2f\t%v\n", fa.balance, fa.money, fa.note)
}
func (fa *familyAccount)exit() {
	for{
		fmt.Print("你确定要退出吗：[y/n]")
		flag := ""
		fmt.Scanln(&flag)
		if flag == "y"{
			fmt.Println("Bye bye!")
			fa.quitFlag = true
			break
		} else if flag == "n"{
			fa.quitFlag = false
			break
		} 
		
	}
}
func (fa *familyAccount) MainMenu()  {
	var opCode int
	for {
		fmt.Println("")
		fmt.Println("--------------家庭收支记账软件-----------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Print("请选择(1-4)：")
		
		fmt.Scanln(&opCode)
		switch opCode{
		case 1:
			fa.showDetial()
		case 2:
			fa.income()
		case 3:
			fa.pay()
		case 4:
			fa.exit()
			if fa.quitFlag {
				return
			}
		default:
			fmt.Println("输入错误，请输入 1-4 之间的数字")
		}
	}
}