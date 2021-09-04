package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TestSwith  case中自带break 不需要写break
func TestSwith(){
	day := time.Now().Weekday().String()
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go")
	default:
		fmt.Println("It's weekend,It's time to rest")
	}
}

// 可以省略switch语句中的条件，
//此模式类似于比较 true 值，就像强制 switch 语句一直运行一样。

//SwichExample 省略switch中的条件
func SwichExample(){
	rand.Seed(time.Now().Unix())
	r := rand.Float64()

	switch  {
	case r> 0.1 :
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}
}

