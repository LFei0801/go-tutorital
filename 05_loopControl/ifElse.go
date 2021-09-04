package main

import "fmt"

func  giveMeANumber() int{
	return -1
}

// TestIfAndElse go中if else 比较奇怪 没有括号
func  TestIfAndElse()  {
	num := giveMeANumber()
	if num < 0{
		fmt.Println("num < 0")
	}else if num == 0{
		fmt.Println("num = 0")
	}else{
		fmt.Println("num > 0")
	}
}

// TestIfAndElse2 if else 支持在if语句中声明变量
func TestIfAndElse2()  {
	if num := giveMeANumber();num < 0{
		fmt.Println("num < 0")
	}else if num == 0{
		fmt.Println("num = 0")
	}else{
		fmt.Println("num > 0")
	}
	//Unresolved reference 'num' 不能在if/else外使用该变量
	//fmt.Println(num)
}

//不支持三元表达式
func testIfAndElse3() {
	num := giveMeANumber()
	// 不支持三元表达式
	//ret :=  num > 0?"num > 0":"num <= 0"
	println(num)
}
