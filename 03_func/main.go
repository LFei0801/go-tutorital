package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(num1 string,num2 string) int {
	int1,_ := strconv.Atoi(num1)
	int2,_ := strconv.Atoi(num2)
	return int1 + int2
}

// 为函数的返回值设置名称，将其当作一个变量
func mul(num1 string,num2 string) (ret int){
	int1,_ := strconv.Atoi(num1)
	int2,_ := strconv.Atoi(num2)
	ret = int1 * int2
	return ret
}

// 返回多个值，返回值用逗号隔开
func calc(num1 string,num2 string) (int , int){
	int1,_ := strconv.Atoi(num1)
	int2,_ := strconv.Atoi(num2)
	return int1+int2,int1 * int2
}

func updateName(name string) {
	// 试图修改name的值
	name = "new Name"
}

func  updateNameByPoint(name *string)  {
	*name = "new Name"
}

func main() {
	// os.Args 变量包含传递给程序的每个命令行参数，值的类型为 string
	result := sum(os.Args[1],os.Args[2])
	mulRet := mul(os.Args[1],os.Args[2])
	sum1,mul1 := calc(os.Args[1],os.Args[2])

	//命令行输入： go run sum.go 2 4
	fmt.Println(result,mulRet,sum1,mul1) //6 8 6 8

	name := "Jhon"
	updateName(name)
	println("now name is: ", name) //now name is:  Jhon
	updateNameByPoint(&name)
	println("now name is: ", name) //now name is:  new Name
}