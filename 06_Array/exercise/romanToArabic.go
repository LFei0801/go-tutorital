package main

import "fmt"

/**
创建罗马数字转换器
使用映射加载要用于将字符串字符转换为数字的基本罗马数字
例如，M 将是映射中的键，其值将为 1000。 使用以下字符串字符映射表列表：

M => 1000
D => 500
C => 100
L => 50
X => 10
V => 5
I => 1

例如，数字 MCM 应打印为 1,900
*/

func romanToArabic(roman string) int{
	// 罗马数字 与 阿拉伯数字的映射
	romanMap := map[string]int{
		"M":1000,
		"D":500,
		"C":100,
		"L":50,
		"X":10,
		"V":5,
		"I":1,
	}
	// 对应的阿拉伯数字集合
	// 为了后续转换规则的进行,集合长度设为 罗马字符串长度 + 1
	arabicList := make([]int,len(roman)+1)

	// 返回的阿拉伯数字
	total := 0

	// 遍历罗马字符串,将每个字母对应的数字存储进阿拉伯集合中
	for index,char := range roman{
		// char 为 int类型,应将其转换为 string类型
		num,exit := romanMap[string(char)]
		if exit{
			arabicList[index ] = num
		}else{
			fmt.Println("输入的罗马数字格式不对")
			return 0
		}
	}

	// 遍历数字集合,如果后一项的数字比前一项大,那么将前一项数字设为对应的负数
	for i:=0;i<len(arabicList)-1;i++{
		if arabicList[i] < arabicList[i+1]{
			arabicList[i] = -arabicList[i]
		}
		total += arabicList[i]
	}
	return total
}
