package main

import "fmt"

func main() {
	/*
		fibonacci 测试
	*/
	fib := fibonacci(6)
	// [1 1 2 3 5 8]
	fmt.Printf("%v\n", fib)
	fib2 := fibonacci(7)
	// [1 1 2 3 5 8 13]
	fmt.Printf("%v\n", fib2)

	/**
	  罗马数字转整数测试
	*/
	// 1900
	fmt.Println("romanToArabic:", romanToArabic("MCM"))
	// 300
	fmt.Println("romanToArabic:", romanToArabic("MMM"))
	// 2800
	fmt.Println("romanToArabic:", romanToArabic("MCMCM"))
	// 输入的罗马数字格式不对
	fmt.Println("romanToArabic:", romanToArabic("SDd"))

	/**
	twoSum测试
	*/
	nums := []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))

	/**
	 firstUniqueChar测试
	 */
	fmt.Printf("unique char is %c\n",firstUniqueChar("google"))
}
