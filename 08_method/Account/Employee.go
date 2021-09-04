package main

import (
	"errors"
	"fmt"
)

// Employee 员工
type Employee struct {
	Account // 账户名
	Salary  float64 //薪水
	Credits float64 // 账户余额
}

// CreateEmployee 员工账户构造函数
func CreateEmployee(name string,credits,Salary float64) *Employee{
	return &Employee{Account{name},Salary,credits}
}

func (e *Employee) String() string{
	return fmt.Sprintf("姓名: %s  工资： %.2f  账户余额：%.2f\n",e.name,e.Salary,e.Credits)
}

//AddCredits 增加账户余额
func (e *Employee) AddCredits() error{
	if e.Salary > 0.0{
		e.Credits += e.Salary
		return nil
	}
	return errors.New("员工的薪水不能为负数")
}

// RemoveCredits  减少账户余额
func (e *Employee) RemoveCredits(account float64) error{
	if e.Credits > account{
		if account > 0.0{
			e.Credits -= account
			return nil
		}
		return errors.New("输入金钱数不能为负数")
	}
	return errors.New("账户余额不够")
}

// CheckSalary 检查员工薪水
func (e *Employee) CheckSalary() float64{
	return e.Salary
}

