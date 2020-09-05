package main

import(
	"fmt"
)

func main()  {
	//账户余额
	balance := 10000.0
	//收支金额
	money := 0.0
	//收支明细
	note := ""
	//收支明细
	detial := "收支\t账户余额\t收支金额\t说  明\n"
	//是否有收支行为
	hasUsed := false
	for {
		fmt.Println("")
		fmt.Println("--------------家庭收支记账软件-----------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Print("请选择(1-4)：")
		
		var opCode int
		fmt.Scanln(&opCode)
		
		switch opCode{
		case 1:
			fmt.Println("--------------当前收支明细记录-----------------")
			if !hasUsed{
				fmt.Println("当前还没收支明细....记一笔吧！")
			} else{
				fmt.Println(detial)
			}
			fmt.Println("")
		case 2:
			fmt.Print("本次收入金额：")
			fmt.Scanln(&money)
			fmt.Print("本次收入说明：")
			fmt.Scanln(&note)
			hasUsed = true
			balance += money
			detial += fmt.Sprintf("收入\t%8.2f\t%8.2f\t%v\n", balance, money, note)
		case 3:
			fmt.Print("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance{
				fmt.Println("支出金额超过余额！")
				continue
			}
			fmt.Print("本次支出说明：")
			fmt.Scanln(&note)
			hasUsed = true
			balance -= money
			detial += fmt.Sprintf("支出\t%8.2f\t%8.2f\t%v\n", balance, money, note)
		case 4:
			fmt.Println("退出软件")
			break;
		default:
			fmt.Println("输入错误")
		}
		if opCode == 4{
			quitFlag := false
			for{
				fmt.Print("你确定要退出吗：[y/n]")
				flag := ""
				fmt.Scanln(&flag)
				if flag == "y"{
					fmt.Println("Bye bye!")
					quitFlag = true
					break
				} else if flag == "n"{
					break
				} 
				
			}
			if quitFlag{
				break;
			} 
		}
	}
}