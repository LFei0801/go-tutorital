package main

import "fmt"

func main() {
	t := triangle{3}
	s := square{3}
	c := colorTriangle{
		triangle{4},
		"red",
	}
	// t perimeter: 9
	fmt.Println("t perimeter:",t.perimeter())
	t.doubleSize()
	// t size: 6
	fmt.Println("t size:",t.size)
	// s perimeter: 12
	fmt.Println("s perimeter:",s.perimeter())

	// c size 4
	fmt.Printf("c size %d\n" ,c.triangle.size)
	// c perimeter(重载后的方法): 24
	fmt.Println("c perimeter(重载后的方法):",c.perimeter())
	// c perimeter(父类方法)： 12
	fmt.Println("c perimeter(父类方法)：",c.triangle.perimeter())

	c1 := circle{2.0}
	// c1 perimeter is : 12.56
	fmt.Println("c1 perimeter is :",c1.perimeter())
	// c1 area is : 12.56
	fmt.Println("c1 area is :",c1.area())

	// p is 12.560000	,s is 12.560000
	printInformation(&c1)

	num := 1
	str := "string"
	f := 1.23
	arr := [...]int{1,2,3,4}
	// e type is int
	printType(num)
	// e type is string
	printType(str)
	// e type is float64
	printType(f)
	// e type is [4]int
	printType(arr)
}
