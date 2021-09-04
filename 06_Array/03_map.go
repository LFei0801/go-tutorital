package main

import "fmt"

func mapBasis()  {
	// 使用make函数创建空映射
	person := make(map[string]int)
	fmt.Println("person:",person) //person: map[]
	// 空映射可以直接添加键值对
	person["Jhon"] = 12
	person["Jack"] = 14
	person["Alice"] = 15
	person["Bob"] = 18
	//	使用范围遍历map中元素
	//  key value
	for name,age := range person{
		fmt.Printf("name is %s , age is %d\n",name,age)
	}
	// nil map
	var nilMap map[int]int
	fmt.Println("nilmap :", nilMap) //nilmap : map[]
	// 	// Assignment to an entry might panic because of the 'nil'
	//nilMap[1] = 1
}

// 获取某一项的值
func getValueBykey(){
	p := make(map[string]int)
	list := []string{"Jack","Alice","Bob"}
	// 赋值
	for index,name := range list{
		p[name] = index
	}
	for name,age := range p{
		fmt.Printf("%s\t%d\t\n",name,age)
	}
	// Nero 不存在，但是不会报错
	fmt.Printf("%v\n",p["Nero"]) //0

	name,exit := p["Nero"]
	if exit{
		fmt.Println("exit , name is :",name)
	}else{
		fmt.Println("Nero not found ~~~ ")
	}

	deleteMap(p,"Alice")
	// current p in main func: map[Bob:2 Jack:0]
	fmt.Println("current p in main func:",p)
}

func  deleteMap(source map[string]int,key string)  {
	// before: map[Alice:1 Bob:2 Jack:0]
	fmt.Println("before:",source)
	delete(source,key)
	// after: map[Bob:2 Jack:0]
	fmt.Println("after:",source)
}
