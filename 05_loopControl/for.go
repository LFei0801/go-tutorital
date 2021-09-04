package main

import (
	"fmt"
)

// 判断num是否是素数
func isPrime(num int) bool{
	for i:=2; i<num; i++ {
		if num % i == 0{
			return  false
		}
	}
	return true
}

// GetPrime 获取 num 范围内所有的素数
func  GetPrime(num int) {
	for i:=3; i<=num; i++{
		if isPrime(i){
			fmt.Printf("%d is prime\n",i)
		}
	}
}

// SimulateWhile  模拟while循环
func SimulateWhile(){
	num := 1
	for num <= 5{
		fmt.Printf("%d\n",num)
		num ++
	}
}

func menuTips ()  {
	fmt.Println("输入1，进入目录1")
	fmt.Println("输入2，进入目录2")
	fmt.Println("输入3，进入目录3")
	fmt.Println("输入4，退出系统")
}

// UnStopLoop 模拟无限循环
func UnStopLoop() {
	menuTips()
	fmt.Println("请输入数字：")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
	}
	for {
		if num == 1{
			fmt.Println("目录1")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}else if num == 2{
			fmt.Println("目录2")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}else if num == 3 {
			fmt.Println("目录3")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}else if num == 4{
			break
		}else{
			fmt.Println("重新输入，谢谢！")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
