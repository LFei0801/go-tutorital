package main

import (
	"fmt"
	"strconv"
)

/*
	声明变量和常量
 */

// 声明变量
//var firstName string = "John"
//var lastName string = "Doe"
//
//// 也可以使用块级来声明变量
//var(
//	first string = "first"
//	last string = "last"
//	//	单行声明多个变量
//	//one,two,three = "one","two","three"
//)
//
//// 也可以不显示声明变量类型，Go会自动根据变量初始值来推断变量类型
//var (
//	age = 12
//	name = "Jhon"
//)



func main() {
	// 常用的方法时是是使用短声明的方式来声明变量
	// 使用冒号等于号时，要声明的变量必须是新变量
	firstName, lastName := "John", "Doe"
	fmt.Println(firstName,lastName) //John Doe

	// 字符串转为数字
	i,_ := strconv.Atoi("-42")
	println(i) //-42
}