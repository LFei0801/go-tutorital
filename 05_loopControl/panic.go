package main

import "fmt"

/**
 执行 PanicExample(2,0) 输出结果：
  Call:high: 2 low: 0
  Call:high: 2 low: 1
  Call:high: 2 low: 2
  panic
  Defered: high: 2 low: 2
  Defered: high: 2 low: 1
  Defered: high: 2 low: 0
 */

func PanicExample(high int,low int){
	if high < low{
		fmt.Println("panic")
		panic(" low greater than high")
	}
	defer  fmt.Println("Defered: high:",high,"low:",low)
	fmt.Println("Call:high:",high,"low:",low)

	PanicExample(high,low + 1)
}
