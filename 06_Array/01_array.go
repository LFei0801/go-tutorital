package main

import "fmt"

/*
 ArrayBasis 基础知识
 len(a) 表示数组的长度
 未被初始化的数组元素会被默认赋予该类型默认值
 */
func ArrayBasis ()  {
	var a[3]int
	a[0] = 10
	a[2] = 11
	fmt.Println(a[0],a[1],a[len(a)-1]) //10 0 11

   // 也可以在声明时初始化
   b := [3]int{1,2,3}
   for i:=0;i<len(b);i++{
   	// 1	2	3
   	fmt.Printf("%d\t",b[i])
   }
   fmt.Printf("\n")

   //  如果你不知道你将需要多少个位置，但知道你将具有多少数据
   //  可以使用省略号省略数组 长度
   c := [...]string{"Wuhan","Chongqing","Beijing"}
   for i := 0;i<len(c);i++{
   	//Wuhan	Chongqing	Beijing
   	fmt.Printf("%s\t",c[i])
   }
   fmt.Printf("\n")

  // 可以为数组的指定位置指定初始值
  d := [...]int{5:-1}
  for i := 0;i<len(d);i++{
  	// 0	0	0	0	0	-1
  	fmt.Printf("%d\t",d[i])
  }
  fmt.Printf("\n")
}
