package main

import "fmt"

/*
 声明 人 数据类型
 */

type person struct {
	ID int
	FirstName string
	LastName string
	Address string
}

/*
 结构也可以嵌套使用
 声明员工数据类型
 */
type employee struct {
	// 员工首先是个人 因此也具有人的数据
	information person
	ManagerID int
}

/**
 同名字段可以简写
 */
type contractor struct {
	employee
	CompanyID int
}



func structBasis() {
	// 初始化Person
	p := person{
		123,
		"Jhon",
		"James",
		"China",
	}
	// {123 Jhon James China}
	fmt.Println(p)
	//使用点语法访问某个字段
	// Jhon	James
	fmt.Printf("%s\t%s\n", p.FirstName, p.LastName)

	// employee数据初始化
	e := employee{information: p,ManagerID: 123}
	// befor：{123 Jhon James China}
	fmt.Printf("befor：%v\n",e.information)
	changeName(&e)
	// after：{123 newFirstName newLastName China}
	fmt.Printf("after：%v\n",e.information)

	// contractor 数据初始化
	var c contractor
	c.employee = e
	c.CompanyID = 123123
	// {{{123 newFirstName newLastName China} 123} 123123}
	fmt.Println(c)
}

func changeName(e *employee){
	e.information.FirstName = "newFirstName"
	e.information.LastName = "newLastName"
}


