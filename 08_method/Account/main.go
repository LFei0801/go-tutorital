package main

import (
	"fmt"
	"os"
)

func main() {
	c := &Company{}
	c.Employees = make(map[int]*Employee)
	c.Credits = 1000.0
	defaultId := 0
	for {
		showFirstLevelMenu()
		fmt.Print("输入：")
		var input int
		_, _ = fmt.Scanf("%d", &input)
		fmt.Println("input is :",input)
		switch input {
		case 1:
			c.CheckAllEmployee()
		case 2:
			err := c.AddEmployee(defaultId)
			if err != nil {
				fmt.Println(err)
			}
			defaultId++
		case 3:
			var id int
			fmt.Print("请输入待删除员工的id:")
			_, _ = fmt.Scanf("%d", &id)
			err := c.RemoveEmployee(id)
			if err != nil {
				fmt.Println(err)
			}
		case 4:
			err := c.PayEmployees()
			if err != nil {
				fmt.Println(err)
			}
		case 5:
			fmt.Printf("请输入查询的员工姓名：")
			var name string
			_, _ = fmt.Scanf("%s", &name)
			e, err := c.findEmployee(name)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(e.String())
		case 6:
			fmt.Print("输入需扣除工资的员工id:")
			var id int
			_,_ = fmt.Scanf("%d",&id)
			fmt.Print("输入需扣除的工资数额：")
			var salary float64
			_,_ = fmt.Scanf("%f",&salary)
			err := c.SalaryDeductionById(id,salary)
			if err != nil{
				fmt.Println("err:",err)
			}
		case 7:
			fmt.Println(c.CheckCompanyCredits())
		case 0:
			os.Exit(0)
		default:
			fmt.Println("输入错误，请重试！")
		}
	}
}
