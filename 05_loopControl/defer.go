package main

import "fmt"

/**
输出结果：
    regular 1
	regular 2
	regular 3
	defered: -3
	defered: -2
	defered: -1

Process finished with the exit code 0

*/

func DeferExample(){
	for i:=1;i<4;i++{
		defer fmt.Println("defered:",-i)
		fmt.Println("regular",i)
	}
}