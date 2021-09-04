package main

import "fmt"

/**
编写一个程序来计算斐波纳契数列
编写一个函数，它返回一个包含按斐波纳契数列排列的所有数字的切片，
而这些数字是通过根据用户输入的大于 2 的数字计算得到的。
小于 2 的数字会导致错误，并返回一个 nil 切片。

斐波纳契数列是一个数字列表，其中每个数字是前两项斐波那契数字之和。
*/

func fibonacci(num int) []int {
	if num < 2 {
		fmt.Println("输入的数字应大于2")
		return make([]int, 0)
	}

	nums := make([]int, num)
	nums[0] = 1
	nums[1] = 1

	for i := 2; i < num; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}

	return nums
}
