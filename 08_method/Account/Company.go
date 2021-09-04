package main

import (
	"errors"
	"fmt"
)

//Company 公司
type Company struct {
	Employees map[int]*Employee // 员工列表
	Credits float64      // 公司所有余额
}

//CheckAllEmployee 查看员工列表
func (c *Company) CheckAllEmployee(){
	if len(c.Employees) == 0{
		fmt.Println("公司暂无员工，请添加员工.....")
		return
	}
	for id,e := range c.Employees{
		fmt.Printf("%d:%s",id,e.String())
	}
	fmt.Println("员工列表输出完毕")
}

//ScanEmployeeInformation 录入员工的信息
func ScanEmployeeInformation() ( string,float64,float64){
	var name string
	fmt.Print("输入员工姓名：")
	_,_ = fmt.Scanf("%s",&name)
	var salary float64
	fmt.Print("输入员工工资（请精确到小数点前两位）：")
	_,_ = fmt.Scanf("%f",&salary)
	var credits float64
	fmt.Print("输入员工账户余额（请精确到小数点前两位）：")
	_,_ = fmt.Scanf("%f",&credits)
	return name,salary,credits
}

// AddEmployee 添加员工
func (c *Company) AddEmployee(id int) error{
	_,exit :=  c.Employees[id]
	if exit{
		return errors.New("添加员工失败,员工已经存在系统中....\n")
	}
	name,salary,credits := ScanEmployeeInformation()
	c.Employees[id] = CreateEmployee(name,credits,salary)
	fmt.Println("保存成功")
	return nil
}

// RemoveEmployee 删除员工
func (c *Company) RemoveEmployee(id int) error{
	_,exit := c.Employees[id]
	if exit{
		delete(c.Employees,id)
		fmt.Println("删除员工成功")
		return nil
	}
	return  errors.New("删除员工成功,员工不存在系统中....\n")
}

// 查询员工
func (c *Company) findEmployee(name string) (*Employee,error){
	for _,e := range c.Employees{
		if e.name == name{
			return e,nil
		}
	}
	return nil,errors.New("查询员工失败,未查询到匹配员工")
}


// PayEmployees 给所有员工发放工资
func (c *Company) PayEmployees() error{
	for _,employee := range c.Employees{
		if c.Credits < employee.CheckSalary(){
			fmt.Printf("从%s起工资没有发放了~~~~\n",employee.name)
			return errors.New("对不起，公司的金钱不足以支付员工工资了....\n")
		}
		c.Credits -= employee.Salary
		_ = employee.AddCredits()
	}
	fmt.Println("所有员工工资已经发放完毕了....")
	return nil
}

//SalaryDeductionById 扣除某位员工工资
func (c *Company) SalaryDeductionById(id int,account float64) error{
	e,exit := c.Employees[id]
	if exit{
		e.Salary -= account
		fmt.Printf("从%s的工资中扣除了%.2f,其本月应发放工资位:%.2f\n",e.name,account,e.Salary)
		return nil
	}
	return errors.New("扣除员工工资失败，员工不存在系统中....\n")
}

//ChangeEmployeeName 更改员工姓名
func (c *Company) ChangeEmployeeName(id int,firstName string) error{
	e,exit := c.Employees[id]
	if exit{
		e.ChangeName(firstName)
	}
	return  errors.New("员工不存在系统中....\n")
}

//CheckCompanyCredits 检查公司账户余额
func (c *Company) CheckCompanyCredits() float64{
	return c.Credits
}
