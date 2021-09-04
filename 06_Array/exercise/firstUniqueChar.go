package main

import "fmt"

/**
找出字符串中第一个未重复的数字
*/

func firstUniqueChar(s string) byte {
	// 哈希表 存储值和下标的映射
	hashTable := make(map[rune]int)
	// 遍历字符串，如果其字符已经存在与哈希表中就将其下标设置-1
	for index, char := range s {
		_, exit := hashTable[char]
		if exit {
			hashTable[char] = -1
		} else {
			// 不存在就存储其值和下标的映射
			hashTable[char] = index
		}

	}
	fmt.Println("hashTable is: ", hashTable)
	first := len(s)
	var ret byte = ' '
	// 遍历哈希表，找到值不为-1的最小值
	for char, pos := range hashTable {
		if pos != -1 && pos < first {
			first = pos
			ret = byte(char)
		}
	}
	return ret
}
