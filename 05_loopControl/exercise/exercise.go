package main

import "fmt"

/*
编写 FizzBuzz 程序
	首先，编写一个用于输出数字（1 到 100）的程序，其中有以下变化：
	如果数字可被 3 整除，则输出 Fizz。
	如果数字可被 5 整除，则输出 Buzz。
	如果数字可同时被 3 和 5 整除，则输出 FizzBuzz。
	如果前面的情况都不符合，则输出该数字。
	尝试使用 switch 语句。
*/

func FizzBuzz(){
	for i:=1;i<=100;i++{
		switch  {
		case i % 15 == 0:
			fmt.Println("FizzBuzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println("i: ",i)
		}
	}
}



/**
查找质数
 编写一个程序来查找小于 20 的所有质数
 */

func FindPrimes(number int) bool {
	for i := 2;i < number;i++{
		if number % 2 == 0{
			return false
		}
	}
	if number > 1{
		return true
	}else{
		return false
	}
}

/**
要求用户输入一个数字，如果该数字为负数，则进入紧急状态
持续要求用户输入一个整数。 此循环的退出条件应该是用户输入了一个负数。
当用户输入负数时，让程序崩溃。 然后输出堆栈跟踪错误。
如果数字为 0，则输出“0 is neither negative nor positive”。 继续要求用户输入数字。
如果数字为正数，则输出“You entered: X”（其中的 X 为输入的数字）。 继续要求用户输入数字。
 */

func ScanNumber(){
	number := 0
	for{
		fmt.Println("Enter number :")
		_,err  := fmt.Scanf("%d",&number)
		if err != nil{
			fmt.Println(err)
		}
		switch{
		case number < 0 :
			panic("You enter a negative number")
		case number == 0:
			fmt.Println("0 is neither negative nor positive")
			ScanNumber()
		default:
			fmt.Println("You Entered: ",number)
			ScanNumber()
		}
	}
}
