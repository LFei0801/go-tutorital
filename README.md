# golang 学习

# Hello World
```go
package main

// import 语句可以访问其他包的代码
import "fmt"

// main 函数是程序的起点,只能有一个main函数
func main()  {
	// 调用fmt包中的Println函数
	fmt.Println("Hello World")
}
```



# 在Go中使用包、变量和函数

学习目标：
- 声明变量和常量
- 了解Go提供的基本数据类型
- 编写函数
- 创建和使用包

## 声明变量和常量
Golang中声明变量的一种方式是使用var关键字

```go
package main

// 声明变量可以显示的给出变量初始类型
var firstname string 

// 初始化变量时只需要在后面给出初始值
var lastName string = "Doe"

// 可以不显示的给出变量类型，go会自动推断变量类型
// 可以在() 内一次性声明多个变量
var (
	one = "one"
	two = "two"
)

// 可以在单行一次性声明多个变量，用逗号隔开
var three,four = "three",4
```

**声明变量最常用的方式是使用 := 方式声明变量**

```go
package  main

func  main()  {
    firstname,lastname := "Jhon","Doe"
    println(firstname,lastname)
    
}
```
使用冒号等于号时，要声明的变量必须是新变量<br>
请记住，对于在 Go 中声明的每个变量，你必须将其用于某处。否则会报错
```
控制台输出报错：
./main.go:4:2: firstName declared but not used
./main.go:4:13: lastName declared but not used
./main.go:5:2: age declared but not used
```

Golang 使用 const关键字声明常量
```go
package main

const LEARN_GOLANG = "learn-golang"
```

## 类型系统
Golang中内置的有四类数据类型：
- 基本类型： 数字、字符串和布尔值
- 聚合类型： 数组和结构
- 引用类型： 指针、切片、映射、函数和通道
- 接口类型: 接口

类型转换：在 Go 中隐式强制转换不起作用，需要显式强制转换，可以使用strconv包中的函数了进行转换：
```go
package main

import "strconv"

func main() {
	// 字符串转为数字
	i,_ := strconv.Atoi("-42")
	println(i) //-42
}
```

**其他复杂数据类型在后面介绍**

## 函数
在 Go 中，函数允许你将一组可以从应用程序的其他部分调用的语句组合在一起，使用func关键字声明函数

### 理解main函数
main函数是整个程序的入口函数，在一个包中只能运行有且仅有一个main函数<br>
如果创建的是 Go 包，则无需编写 main() 函数<br>

### 声明函数 和 使用函数
语法：
```
// 函数声明
func 函数名 (参数 参数类型，参数 参数类型) 返回值类型 {
	函数逻辑
}

//  函数调用
函数名（参数，参数 ...）
```

Go函数的特点：

- 可以给返回值命名，当一个变量使用
- 可以返回多个值
- 使用 _ 来忽略某个函数返回值
- 当两个参数的类型一样是可以省略第一个参数类型

当返回多个值时，语法为：
```
// 返回值名称可以省略
func 函数名（变量名 变量类型，变量名 变量类型...） (返回值名 返回值类型，返回名 返回类型...)
```

比如 求命令传入的两个参数之和：
```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(num1 string,num2 string) int {
	int1,_ := strconv.Atoi(num1)
	int2,_ := strconv.Atoi(num2)
	return int1 + int2
}

// 为函数的返回值设置名称，将其当作一个变量
func mul(num1 string,num2 string) (ret int){
	int1,_ := strconv.Atoi(num1)
	int2,_ := strconv.Atoi(num2)
	ret = int1 * int2
	return ret
}

// 返回多个值，返回值用逗号隔开
func calc(num1 string,num2 string) (int , int){
	int1,_ := strconv.Atoi(num1)
	int2,_ := strconv.Atoi(num2)
	return int1+int2,int1 * int2
}

func main() {
	// os.Args 变量包含传递给程序的每个命令行参数，值的类型为 string
	result := sum(os.Args[1],os.Args[2])
	mulRet := mul(os.Args[1],os.Args[2])
	sum1,mul1 := calc(os.Args[1],os.Args[2])
	//命令行输入： go run sum.go 2 4
	fmt.Println(result,mulRet,sum1,mul1) //6 8 6 8
}
```

### 更改函数参数值
将值传递给函数时，该函数中的每个更改都不会影响调用方。<br>
Go 是“按值传递”编程语言。 这意味着每次向函数传递值时, Go 都会使用该值并创建本地副本（内存中的新变量）。
在函数中对该变量所做的更改都不会影响你向函数发送的更改

可以看这个例子：
```go
package main

func updateName(name string){
	name = "new name"
}

func main() {
    name := "Jhon"
	updateName(name)
    println("now name is: ",name) // now name is Jhon
}
```
可以发现name的值并未修改，这是因为name是将其值传给updateName函数，函数调用时，会在内存中新开辟一个地址
储存name的副本，修改的是这个副本的值，并未影响原本的值。

**解决思路： 使用指针传递name的地址**
指针 是包含另一个变量的内存地址的变量。 当你发送指向某个函数的指针时，不会传递值，而是传递地址内存
因此，对该变量所做的每个更改都会影响调用方。

在 Go 中，有两个运算符可用于处理指针

- & 运算符 使用其对象的地址
- * 运算符 前往指针中包含的地址访问其中的对象

```go
package main

func updateName(name string){
	name = "new name"
}

// *string 表示name为指向字符串类型的指针
func updateNameByPoint(name *string){
	//*name 取出name指针指向地址的值
	*name = "new Name"
}

func main() {
    name := "Jhon"
	updateName(name)
    println("now name is: ",name) // now name is Jhon
    // &name 表示传递name的地址
	updateNameByPoint(&name)
    println("now name is: ",name) // now name is new Name
}
```

## 包
Go 包与其他编程语言中的库或模块类似。 你可以打包代码，并在其他位置重复使用它。 
包的源代码可以分布在多个 .go 文件中

### 创建自定义包的规则
- 如需将某些内容设为专用内容，请以小写字母开始
- 如需将某些内容设为公共内容，请以大写字母开始

```go
package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

// 只能在包内调用该函数
func internalSum(num int) int{
	return num - 1
}

// Sum two int numbers
func Sum(num1 ,num2 int) int{
	return num1 + num2
}
```

### 使用自定义包
新建usePackage文件，新建main.go文件，项目目录如下：
```
src/
  calculator/
    sum.go
  usePackage/
    main.go
```
在main.go文件中使用calculator包的公共函数和变量：
```
package main

import "../calculator"

func main() {
	total := calculator.Sum(3,5)
	println(total) //8
	println("Version:",calculator.Version) //Version: 1.0
}
```

### 使用包的规则
包名称遵循的约定：
- 在 Go 中，包名称需遵循约定。 包使用其导入路径的最后一部分作为名称

```
import "math/cmplx"
```
若要引用包中的对象，请使用包名称cmplx
```
cmplx.Inf()
```

# 在Go中使用控制流
学习目标：
- 了解简单的和复合的if语句
- 了解switch 语句和基本功能
- 使用for关键字来了解loop语句及其在 Go 中的模式
- 使用基本函数来处理 defer、panic 和 recover 之类的错误

## 使用if/else表达式测试条件
Golang中if/else 特点：
- 在 Go 中，你不需要在条件中使用括号。else 子句可选。但是，大括号仍然是必需的
- 不支持三元表达式
- 可以在 if 条件中声明变量，其作用域仅限 if 块

```go
package main

import "fmt"

func  giveMeANumber() int{
	return -1
}

// go中if else 比较奇怪 没有括号
func  testIfAndElse()  {
	num := giveMeANumber()
	if num < 0{
		fmt.Println("num < 0")
	}else if num == 0{
		fmt.Println("num = 0")
	}else{
		fmt.Println("num > 0")
	}
}

// if else 支持在if语句中声明变量
func testIfAndElse2()  {
	if num := giveMeANumber();num < 0{
		fmt.Println("num < 0")
	}else if num == 0{
		fmt.Println("num = 0")
	}else{
		fmt.Println("num > 0")
	}
	//Unresolved reference 'num' 不能在if/else外使用该变量
	//fmt.Println(num)
}

//不支持三元表达式
func testIfAndElse3() {
	num := giveMeANumber()
	// 不支持三元表达式
	//ret :=  num > 0?"num > 0":"num <= 0"
	println(num)
}

func main() {
	testIfAndElse2()
}

```

## switch语法
像if语句一样，switch条件不需要括号。
case中自带break
```go
package main

import (
	"fmt"
	"time"
	"math/rand"
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

//SwichExample 省略switch中的条件
func SwichExample(){
	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	// 程序始终运行这种类型的 switch 语句，因为条件始终为 true
	switch  {
	case r> 0.1 :
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}
}
```

## for循环遍历数据
与 if 语句和 switch 语句一样，for 循环表达式不需要括号。 但是，大括号是必需的
分号 (;) 分隔 for 循环的三个组件：
- 在第一次迭代之前执行的初始语句（可选）
- 在每次迭代之前计算的条件表达式。 该条件为 false 时，循环会停止
- 在每次迭代结束时执行的后处理语句（可选）

基本使用：比如获取范围内的所有素数：
```go
package main

import (
	"fmt"
)

// 判断num是否是素数
func isPrime(num int) bool{
	for i:=2; i<num; i++ {
		if num % i == 0{
			return  false
		}
	}
	return true
}

// GetPrime 获取 num 范围内所有的素数
func  GetPrime(num int) {
	for i:=3; i<=num; i++{
		if isPrime(i){
			fmt.Printf("%d is prime\n",i)
		}
	}
}
```

### 模拟while语句
Golang中没有while关键字，因此必须使用for来模拟while
因为for三个组件都是可选的，因此可以省略第一个和第三个组件来模拟while语句
```go
package main

import "fmt"

// SimulateWhile  模拟while循环
func SimulateWhile(){
	num := 1
	for num <= 5{
		fmt.Printf("%d\n",num)
		num ++
	}
}
```

### 无限循环
不编写条件表达式，也不编写预处理语句或后处理语句，而是采取退出循环的方式进行编写
break和continue关键字同其他语言作用一样，就不再演示

可以使用无限循环模拟系统操作：
```go
package main

import (
	"fmt"
)

func menuTips ()  {
	fmt.Println("输入1，进入目录1")
	fmt.Println("输入2，进入目录2")
	fmt.Println("输入3，进入目录3")
	fmt.Println("输入4，退出系统")
}

// UnStopLoop 模拟无限循环
func UnStopLoop() {
	menuTips()
	fmt.Println("请输入数字：")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println(err)
	}
	for {
		if num == 1{
			fmt.Println("目录1")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}else if num == 2{
			fmt.Println("目录2")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}else if num == 3 {
			fmt.Println("目录3")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}else if num == 4{
			break
		}else{
			fmt.Println("重新输入，谢谢！")
			fmt.Println("请重新输入数字：")
			_, err := fmt.Scanf("%d", &num)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

```

## 使用defer、panic 和 recover 函数进行控制

### defer
在 Go 中，defer 语句会推迟函数（包括任何参数）的运行，直到包含 defer 语句的函数完成
相当于在函数调用栈中，被defer标记的函数会优先入栈
```go
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
```
程序分析：根据defer语句的特性得出的该函数的调用栈为：
fmt.Println("defered:",-1) -> fmt.Println("defered:",-2) -> fmt.Println("defered:",-3)
-> fmt.Println("defered:",3) -> fmt.Println("defered:",2) -> fmt.Println("defered:",1)
根据栈的特性，先进后出，因此该函数的输出为 1，2，3，-3，-2，-1

**通常情况下，当你想要避免忘记任务（例如关闭文件或运行清理进程）时，可以推迟某个函数的运行**

### panic函数和recover函数
运行时错误会使 Go 程序崩溃，例如尝试通过使用超出范围的索引或取消引用 nil 指针来访问数组。 你也可以强制程序崩溃
**panic**函数就是强制令程序崩溃。
调用 panic() 函数时，可以添加任何值作为参数
当你使用 panic 调用时，任何延迟的函数调用都将正常运行。 进程会在堆栈中继续，直到所有函数都返回。 然后，程序会崩溃并记录日志消息

```go
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
```
根据输出结果显示：defer标记的函数在panic后依然可以运行。
在发生未预料到的严重错误时，系统通常会运行对 panic() 函数的调用。 
若要避免程序崩溃，可以使用名为 recover() 的另一个函数

### recover函数
Go 提供内置 recover() 函数，让你可以在程序崩溃之后重新获得控制权。

# 在Go中输用数据类型以及结构、数组、切片和映射

学习目标：
- Go中的聚合类型： 数组和切片
- 数组和切片之间的区别
- 用于操作数据的内置函数
- 如何通过映射使用键值数据结构
- 如何使用结构编写复杂的自定义类型

## 使用数组
Go 中的数组是一种特定类型且长度固定的数据结构。 它们可具有零个或多个元素，必须在声明或初始化它们时定义大小
此外，它们一旦创建，就无法调整大小。 
鉴于这些原因，数组在 Go 程序中并不常用，但它们是切片和映射的基础

**声明数组**
要在 Go 中声明数组，必须定义其元素的数据类型以及该数组可容纳的元素数目
```go
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
```

## 切片 slice
如果习惯了Javascript的动态类型，那么使用Go的数组是一定不会习惯的，甚至有点难受！
JS中的数组更加类似于栈或者队列这种数据结构，可以动态添加数组元素。
Go中的切片比较类似于JS中的数组，不过由于静态类型语言的限制，切片只能存放一系列类型相同的元素。
与数组一样，切片也是 Go 中的一种数据类型，它表示一系列类型相同的元素。
不过，与数组更重要的区别是切片的大小是动态的，不是固定的。

**切片是数组或另一个切片之上的数据结构。 我们将源数组或切片称为基础数组。 通过切片，可访问整个基础数组，也可仅访问部分元素**
切片只有 3 个组件：
- 指向基础数组中第一个可访问元素的指针。 此元素不一定是数组的第一个元素 array[0]。
- 切片的长度。 切片中的元素数目
- 切片的容量。 切片开头与基础数组结束之间的元素数目。

从切片的组件可以看出，**切片是原始数组的一个子集**

```go
package main

import "fmt"

// 切片与数组的区别不大。 可用相同的方式声明这两者
func sliceBasis()  {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months) //["January",....,"December"]
	fmt.Println(len(months)) //12
	fmt.Println(cap(months)) //12
}
```

### make函数初始化slice
使用 := 初始化slice时往往切片长度会出现问题
```
s1 := []int{1}
//1  1
fmt.Printf("%d\t %d\n",len(s1),cap(s1))
s2 := []int{1,2}
//2	 2
fmt.Printf("%d\t %d\n",len(s2),cap(s2))
```
当要使用s1[2]是会报错
```
// panic: runtime error: index out of range [2] with length 1
fmt.Printf("%d\n",s1[2])
```
但是使用make函数创建的切片时不会出现这个问题
make会直接创建含有n个元素的切片,切片长度和容量都为n
make语法为：
```
make([]数据类型,数据容量)
```
比如:
```go
package main

import "fmt"

func main()  {
	// 创建一个含有10个元素的切片，初始值都为该数据类型的默认值
	s := make([]int,10)
	// [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(s)
	// cap: 10
	fmt.Println("cap:",cap(s))
	// len: 10
	fmt.Println("len:",cap(s))
}
```
### 切片项
Go支持切片运算符s[i:p] 其中
- s 代表原始数组或切片的一个新切片
- 左闭右开  即新切片中的元素为 a[i] ~ a[p-1]
- 切片只能引用元素的子集。

```go
package main

import "fmt"

// 切片与数组的区别不大。 可用相同的方式声明这两者
func sliceBasis()  {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months) //["January",....,"December"]
	fmt.Println(len(months)) //12
	fmt.Println(cap(months)) //12

   // 在months切片上创建四个新切片
   createNewSlice(months)
}

// 切片项促使
func createNewSlice(months []string){
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	// [January February March] 3 12
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	//  [April May June] 3 9
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	// [July August September] 3 6
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	// [October November December] 3 3
	fmt.Println(quarter4, len(quarter4), cap(quarter4))
}
```
可以观察到 四个切片的长度相同，但是容量不同。切片的容量表示在原始数据上可以扩展的元素数量。
quarter1 的第一个元素指向了原始数组的第一个元素，因为原始数组的容量为12个，因此可以扩展的容量为12个
类似的可以看quarter4：
第一个元素指向原始数组的第9个元素，因此容量为 12 - 9
因此 可以得出**切片的容量**为：
*cap = cap(原始数据) - 切片第一个元素指向的位置*

```go
package main

import "fmt"

// 切片项促使
func createNewSlice(months []string){
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	// [January February March] 3 12
	fmt.Println(quarter1, len(quarter1), cap(quarter1))
	//  [April May June] 3 9
	fmt.Println(quarter2, len(quarter2), cap(quarter2))
	// [July August September] 3 6
	fmt.Println(quarter3, len(quarter3), cap(quarter3))
	// [October November December] 3 3
	fmt.Println(quarter4, len(quarter4), cap(quarter4))

	// 测试切片的容量
	getCap(quarter1) //11
	getCap(quarter2) //8
	getCap(quarter3) //5
	getCap(quarter4) //2
}

func getCap(quarters []string){
	newQuarter := quarters[1:2]
	fmt.Println(cap(newQuarter))
}
```

因此切片就是原始数据的子集，子集的容量取决于原始数据和切片的操作。

### 追加项
切片的大小不是固定的，而是动态的
Go 提供了内置函数 append(slice, element)，便于你向切片添加元素
返回值为一个新的切片

```go
package main

import "fmt"

func appendSlice(){
	numbers := []int{}
	for i:=0;i<10;i++ {
        numbers = append(numbers,i)
        fmt.Printf("%d\ncap=%d\t%v\n",i,cap(numbers),numbers)
	}
}
```
输出结果为：
```
0	cap=1	[0]
1	cap=2	[0 1]
2	cap=4	[0 1 2]
3	cap=4	[0 1 2 3]
4	cap=8	[0 1 2 3 4]
5	cap=8	[0 1 2 3 4 5]
6	cap=8	[0 1 2 3 4 5 6]
7	cap=8	[0 1 2 3 4 5 6 7]
8	cap=16	[0 1 2 3 4 5 6 7 8]
9	cap=16	[0 1 2 3 4 5 6 7 8 9]
```
当切片容量不足以容纳更多元素时，Go 的容量将翻倍。
第 3 次迭代，此时容量变为 4，切片中只有 3 个元素。 在第 5 次迭代中，容量又变为 8，第 9 次迭代时变为 16

### 删除项
Go中没有内置函数直接删除切片中的某一项，不过可以利用append函数的特点来达到函数项目的。
append(slice,numbers) slice为切片，numbers为添加到slice中的数据。
因此 如果要删除slice[index] 这一项，可以将slice[index+1] 后面的项添加到slice[:index]这个切片中达到删除目的。

```go
package main

import "fmt"

func deleteSlice(slice []int,index int){
	remove := 2
	if len(slice) > remove{
		fmt.Println("Before: ",slice,"remove:",slice[remove])
		slice = append(slice[:index],slice[index+1:]...)
		fmt.Println("After: ",slice)
	}
}
```

### 复制切片
更改切片中的元素时，基础数组将随之更改, 引用该基础数组的任何其他切片都会受到影响
```go
package main

import "fmt"

// 复制切片
func copySlice(){
	letters := []string{"A","B","C","D","E"}
	//Before [A B C D E]
	fmt.Println("Before",letters)

	slice1 := letters[0:2]
	slice2 := letters[1:4]

	slice1[1] = "Z"

	// After [A Z C D E]
	fmt.Println("After",letters)
	// Slice2 [Z C D]
	fmt.Println("Slice2",slice2)
}
```
此时可以使用make函数新建一个新的切片，这个切片不引用任何原始数据，在使用copy函数将数据复制到新切片中，
这样就避免了切片之间的数据污染

```go
// 复制切片
package main

import "fmt"

func copySlice(){
	letters := []string{"A","B","C","D","E"}
	//Before [A B C D E]
	fmt.Println("Before",letters)

	slice1 := letters[0:2]
	slice2 := make([]string,3)
	copy(slice2,letters[1:4])

	slice1[1] = "Z"

	// After [A Z C D E]
	// 影响了原数组
	fmt.Println("After", letters)
	// Slice2 [B C D]
	// 并未影响复制后的切片
	fmt.Println("Slice2", slice2)
}
```

## 映射 map
Go 中的映射是一个哈希表，是键值对的集合
特点：
- 映射中所有的键都必须具有相同的类型
- 映射中所有的值都必须具有相同的类型
- 键和值类型可以不同

map 基本语法为
```
map[key类型]value类型
比如：
map[string]int

访问map中的元素通过键值查找：
map[key name]
```
声明一个map:
```
使用make函数声明一个空映射
newMap := make(map[string]int)

声明一个nil映射
var nilMap map[string]int
```
注意：make声明的空映射可以直接通过键值对方式添加只，nil映射不行
```go
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
```

### 访问项
访问map中的元素通过map[key] 方式查找对应的value
但是 访问一个不存在的key时，不会抛出错误，而是返回对应value类型的默认值
```go
package main

import "fmt"

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
}
```
其实访问一个key时，会返回两个值：
- value key对应的value，没有则返回该类型的默认值
- exit  存在 则 为true 不存在 则 为false

```go
package main
import "fmt"
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
}
```

### 删除项
通过delele函数删除某一项
```
delete(map,key)
```
```go
package main

import "fmt"

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

```

### slice 和 map 总结
JS里 数据分为引用类型和数值类型，那么在Go里：
数值类型的数据有：
- string
- int
- float
- bool
- struct
- interface
- array
引用类型的数据有：
- slice
- map
- point

## 使用结构
Go 中的结构也是一种数据结构，它可包含零个或多个任意类型的字段，并将它们表示为单个实体
 
### 声明结构
```go
package main

import "fmt"

/*
 声明 人 数据类型
 */

type person struct {
	ID int
	FirstName string
	LastName string
	Address string
}

/*
 结构也可以嵌套使用
 声明员工数据类型
 */
type employee struct {
	// 员工首先是个人 因此也具有人的数据
	information person
	ManagerID int
}

/**
 同名字段可以简写
 */
type contractor struct {
	employee
	CompanyID int
}

func structBasis() {
	// 初始化Person
	p := person{
		123,
		"Jhon",
		"James",
		"China",
	}
	// {123 Jhon James China}
	fmt.Println(p)
	//使用点语法访问某个字段
	// Jhon	James
	fmt.Printf("%s\t%s\n", p.FirstName, p.LastName)

	// employee数据初始化
	e := employee{information: p,ManagerID: 123}
	// {123 Jhon James China}
	fmt.Printf("%v\n",e.information)

	// contractor 数据初始化
	var c contractor
	c.employee = e
	c.CompanyID = 123123
	// {{{123 Jhon James China} 123} 123123}
	fmt.Println(c)
}
```

### 结构体指针
使用结构体指针就可以在函数中改变形参中的值了
```
func changeName(e *employee){
	e.information.FirstName = "newFirstName"
	e.information.LastName = "newLastName"
}
```

### 用JSON编码和解码结构
可使用结构来对 JSON 中的数据进行编码和解码。 Go 对 JSON 格式提供很好的支持，该格式已包含在标准库包中。
你还可执行一些操作，例如重命名结构中字段的名称。 例如，假设你不希望 JSON 输出显示 FirstName 而只显示 name，或者忽略空字段， 
可使用如下例所示的字段标记
```
type Person struct {
    ID        int    
    FirstName string `json:"name"`
    LastName  string
    Address   string `json:"address,omitempty"`
}
```
将结构编码成JSON,使用json.Marshal函数
```
data,err :=  json.Marshal(数据切片)
```
将 JSON 字符串解码为数据结构，使用json.Unmarshal 函数。
```
data,err :=  json.Marshal(数据切片)
var decode 数据切片
err := json.Unmarshal(data,&decode)
```

# 在Go中实现错误处理和日志记录
学习目标：
- 在Go中的错误处理方法
- 错误处理策略
- 用于日志记录的log标准包
- 记录框架

## 错误处理
编写程序时，需要考虑程序失败的各种方式，并且需要管理失败
无需让用户看到冗长而混乱的堆栈跟踪错误。 让他们看到有关错误的有意义的信息更好
panic和recover是程序中的异常和意外行为，和错误不一样
err是已知的失败，程序应该可以处理他们
Go 的错误处理方法只是一种只需要 if 和 return 语句的控制流机制
因为Go中函数可以返回多个返回值，因此可以将其中的一个返回值设为err,再在if中处理错误
```
if employee,err := getInformation(1000)
if err != nil{
    处理错误
}
```
如果错误不是nil，则意味着有错误，需要处理
可以打印该错误消息，也可以记录该消息（更可取）

### 错误处理策略
当函数返回错误时，通常是最后一个返回值，用方负责检查是否存在错误并处理错误。
常见的错误处理是将错误打印出来
```go
package main

import "fmt"

type employee struct {
	id int
	firstname string
	lastname string
	address string
}

func getInformation(id int) (*employee,error){
	e,err := apiCallEmployee(1000)
	if err != nil{
		// 常见的错误处理是将其打印出来
		return nil,fmt.Errorf("Got an error when gettin the information,%v\n",err)
	}
	return e,nil
}

func apiCallEmployee(id int) (*employee,error){
	e := employee{firstname: "Doe",lastname:"Jhon"}
	return &e,nil
}
```
另一种策略是在错误为暂时性错误时运行重试逻辑。 
例如，可以使用重试策略调用函数三次并等待两秒种
```
func getInformation(id int) (*employee,error){
	for tries := 0;tries<3;tries++{
		emp,err := apiCallEmployee(1000)
		if err == nil{
			return emp,nil
		}
		// 有错误就等待两秒钟重新尝试运行
		fmt.Println("Server is not responding,retrying...")
		time.Sleep(time.Second * 2)
	}
	return nil,fmt.Errorf("server has failed to respond to get the employee information")
}
```
### 创建可以重用的错误消息
在 Go 中，你可以使用 errors.New() 函数创建错误并在若干部分中重复使用这些错误，如下所示：
```
var errnotFound = errors.New("employee not found")

func getInformationById(id int) (*employee,error){
	if id != 1000{
		return nil,errnotFound
	}
	emp := employee{lastname: "Doe",firstname: "Jhon"}
	return &emp,nil
}
```
请注意，惯例是为错误变量添加 Err 前缀。
errors.Is() 函数允许比较获得的错误的类型
```
employee, err := getInformation(1000)
if errors.Is(err, ErrNotFound) {
    fmt.Printf("NOT FOUND: %v\n", err)
} else {
    fmt.Print(employee)
}
```
### 用于错误处理的推荐方法
在 Go 中处理错误时，请记住下面一些推荐做法：
- 始终检查是否存在错误，即使预期不存在。 然后正确处理它们，以免向最终用户公开不必要的信息
- 在错误消息中包含一个前缀，以便了解错误的来源。 例如，可以包含包和函数的名称
- 创建尽可能多的可重用错误变量。
- 了解使用返回错误和 panic 之间的差异
- 在记录错误时记录尽可能多的详细信息，并打印出最终用户能够理解的错误。

## 日志记录
用户只需要简单的错误信息，以便理解
但是开发人员需要更见详细的错误信息，以便编写补丁

### log包
log包是Go 提供了一个用于处理日志的简单标准包
基本语法：
```go
package main

import "log"

func main() {
	log.Print("Hey,i'm a log")
}
```
输出为:
```
2021/08/31 13:59:24 Hey,i'm a log
```

使用 log.Fatal() 函数记录错误并结束程序，就像使用 os.Exit(1) 一样
```
func logFatal(){
	log.Fatal("Hey,i'm a log by fatal")
	// 不会运行下面这行代码
	fmt.Print("Can you see me!")
}
```
输出为
```
2021/08/31 14:04:20 Hey,i'm a log by fatal
```
log.panic() 函数会打印出具体的堆栈错误
```
func logPanic(){
	log.Panic("Hey,i'm a log error by panic")
	fmt.Println("Can you see me")
}
```
```
2021/08/31 14:06:31 Hey,i'm a log error by panic
panic: Hey,i'm a log error by panic

goroutine 1 [running]:
log.Panic(0xc00010ff48, 0x1, 0x1)
	D:/Program Files/Go/src/log/log.go:354 +0xb7
main.logPanic()
	F:/Go/src/go-tutortial/07_error/02_log.go:18 +0x68
main.main()
	F:/Go/src/go-tutortial/07_error/main.go:6 +0x27

Process finished with the exit code 2
```
另一重要函数是 log.SetPrefix()。 可使用它向程序的日志消息添加前缀
```
func logSetPrefix(){
	log.SetPrefix("添加的前缀():")
	log.Print("i'm a log ")
	log.Fatal("i'm a fatal log")
	fmt.Println("Can you see me")
}
```
输出：
```
添加的前缀():2021/08/31 14:08:36 i'm a log 
添加的前缀():2021/08/31 14:08:36 i'm a fatal log
```
### 记录到文件中
在文件中添加日志后，可以将所有日志集中在一个位置，并将它们与其他事件关联
```
func logToFile(){
   // 创建日志日志文件并配置权限
	file,err := os.OpenFile("info.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
    // 配置记录到文件
	log.SetOutput(file)
	
	log.Print("Hey,i'm a log")
}
```

# 在Go中使用方法和接口
学习目标：
- Go如何实现oop原则，如封装和组合
- 如何写入方法及使用此类方法的原因
- 如何写入嵌入方法和重载方法
- 如何写入和使用接口，以及接口不同于其他编程语言中接口的原因。

## 在Go中使用方法
Go 中的方法是一种特殊类型的函数，但存在一个简单的区别
**必须在函数名称之前加入一个额外的参数。 此附加参数称为 接收方**
Go 中的这一方法类似于在其他编程语言中创建类，因为它允许你实现面向对象编程 (OOP) 模型中的某些功能，
例如嵌入、重载和封装。

### 声明方法
通过添加方法，可以给自定义结构添加特定的行为，声明方法的语法如下：
```
func (variable type) MethodName(parameters...){}
```
在声明方法之前，必须先创建结构。比如给三角形结构添加计算周长方法
```go
package main

import "fmt"

type triangle struct {
	size int
}

// 计算周长
func (t triangle) perimeter() int{
	return t.size * 3
}

func main() {
	t := triangle{3}
	// perimeter: 9
	fmt.Println("perimeter:",t.perimeter())
}
```
方法类似与其他语言的类，只是Go将类拆散开用结构体和方法来模拟类，比如js中声明三角形这个类：
```js
class Triangle{
    constucter(size){
        this.size = size
    }
    perimeter(){
        return this.size * 3
    }
}
// 调用时
const t = new Triangle(3)
console.log(t.perimeter())
```
对比一下可以发现，go中的函数前面的声明的接收方是一种控制权，
只有拥有这个控制权才能使用这个方法。
如果直接调用这个函数时,系统会报错，显示这个函数没有定义：
```
# go-tutortial/08_method
.\01_method.go:16:2: undefined: perimeter
```
因此，调用方法的唯一方式是声明一个结构，拥有方法的访问权限。
这意味着，只要接受方不同，那么即使函数名一样，那么方法也不一样。
```go
package main

import "fmt"

type triangle struct {
	size int
}

type square struct {
	size int
}

// 计算三角形周长
func (t triangle) perimeter() int{
	return t.size * 3
}

//计算正方形周长
func (s square) perimeter() int{
	return s.size * 4
}

func main() {
	t := triangle{3}
	s := square{3}
	// t perimeter: 9
	fmt.Println("t perimeter:",t.perimeter())
	// s perimeter: 12
	fmt.Println("s perimeter:",s.perimeter())
}
```
### 结构体与方法
通过以上的代码可以总结出以下几点：
- 定义一个结构体，就类似与定义一个类
- 结构体中的不同类型数据，类似与类中的属性
- 为结构体创建一个方法，就类似与类中的方法
- 通过数据名称的首字母大小写，和方法名称首字母的大小写来控制其是否是私有方法和私有属性
- 小写字母开头为私有属性或者私有方法

### 方法中的指针
在 Go 中调用函数时，Go 都会复制每个参数值以便使用。
因此如果方法需要更新变量，或者参数过大时，就需要使用指针传递变量的地址。
比如更新三角形的边长为两倍
```
func (t *triangle) doubleSize(){
    // 即使t为指针结构体，访问属性时依然可以使用.语法
    // 这种是 （*t）.size 的语法糖
	t.size *= 2
}
```
如果方法仅可访问接收方的信息，则不需要在接收方变量中使用指针
但是，依据 Go 的约定，如果结构的任何方法具有指针接收方，则此结构的所有方法都必须具有指针接收方，即使某个方法不需要也是如此
因此，在Go中声明方法，**最好是令接收方为指针，传递其地址**

### 继承
在Go中可以使用嵌套结构体来达到继承的目的
比如定义一个有色彩的三角形结构：
```
type colorTriangle struct{
    triangle
    color string
}
```
此时可以发现colorTriange类型的数据依然可以调用triangle的方法

```go
package main

import "fmt"

type triangle struct {
	size int
}

type colorTriangle struct {
	triangle
	color string
}

// 计算三角形周长
func (t triangle) perimeter() int {
	return t.size * 3
}

func main() {
	c := colorTriangle{
		triangle{4},
		"red",
	}
	// 依然可以调用triangle的方法
	fmt.Println("c perimeter is:",c.perimeter())
}
```
此时triangle就是colorTriangle的基类，colorTriangle从其父类中继承下来了其方法的访问权
达到这个目的，这是因为Go自动为我们创建了如下方法：
```
func (t coloredTriangle) perimeter() int {
    return t.triangle.perimeter()
}
```
注意接受者是coloedTriangle,其调用了内部的triangle的方法，
**但是这者方法开发者不需要创建，Go自动完成了**

### 重载
正如上面所说，方法是一种访问权限，只要接受者不同，函数名可以相同，因此可以手动编写父类的方法，
改变其接收者即可达到重载目的：
```
// 重载周长方法
func (c colorTriangle) perimeter() int{
	return c.size * 3 * 2
}
```
如果没有手动写这个方法，那么Go默认其采用父类的方法，手动写了，那么就调用此方法。
基于结构嵌套的特点，即使方法重载了，依然可以调用父类的原始方法
```go
package main

import "fmt"

type triangle struct {
	size int
}
type colorTriangle struct {
	triangle
	color string
}
func (t *triangle) perimeter() int{
	return t.size * 3
}
func (c *colorTriangle) perimeter() int{
	return c.size * 3 * 2
}

func main() {
    c := colorTriangle{
    	triangle{4},
    	"red",
    }
    fmt.Println("c perimeter(重载后的方法):",c.perimeter())
    fmt.Println("c perimeter(父类方法)：",c.triangle.perimeter())
}
```

### 方法总结
- Go中方法就是一种控制权，配合结构体达到oop的目的。
- oop中的重载、继承、封装都是使用了结构体和方法的特点来模拟实现
- 在Go是没有对象Object这个数据类型，可以使用结构体和方法来模拟对象

## 在Go中使用接口
Go 中的接口是一种用于表示其他类型的行为的数据类型
接口类似于对象应满足的蓝图或协定

### 声明接口
Go 中的接口是一种抽象类型，只包括具体类型必须拥有或实现的方法
比如创建一个几何接口
```
type Shape interface{
    Perimeter() float
    Area() float
}
```
如果一个结构实现了这两个方法，那么就说明该结构是Shape类型数据
```go
package main

type shape interface {
	perimeter() float64
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) perimeter() float64{
	return 2 * 3.14 * c.radius
}

func (c circle) area() float64{
	return 3.14 * c.radius * c.radius
}
```
在上面这个程序中，circle结构实现了shape接口要求的两个方法，那么可以说circle类型的数据就是shape类型的数据。
换句话说，**接口是一个抽象类型的数据，具有某些特定行为的数据**。
只要某个结构实现了接口要求的方法，那么这个结构就是拥有该特定行为的数据。

### 函数与接口
函数的形参必须要求确定形参变量的数据类型，那么一个函数是否可以接受一个接口类型的数据呢？
```
func printInformation(s shape){
	fmt.Printf("%T\n",s)
	fmt.Printf("p is %f\t,s is %f\n",s.perimeter(),s.area())
}
```
printInformation负责打印shape类型的数据，其接受了shape类型的数据形参，
调用此方法：
```
// p is 12.560000	,s is 12.560000
printInformation(c1)
```
c1是circle类型的数据，而printInformation 接受了从c1参数，说明c1就是shape接口类型的数据。
这说明**一切实现了接口方法的数据都是该接口类型的数据**
### 指针与接口
如果修改circle的方法就接收者，该为指针类型的接受者
```
func (c *circle) perimeter() float64{
	return 2 * 3.14 * c.radius
}

func (c *circle) area() float64{
	return 3.14 * c.radius * c.radius
}

// printInformation不变
func printInformation(s shape){
	fmt.Printf("%T\n",s)
	fmt.Printf("p is %f\t,s is %f\n",s.perimeter(),s.area())
}
```
此时调用printInformation函数，就不能传c1了，系统会报错：
```
	printInformation(c1)
   // 错误提示：
   //Type does not implement 'shape' as the 'perimeter' method has a pointer receiver
```
错误提示显示c1不是shape类型的数据，因为其方法是指针类型的接收者。
那么是否意味着c1的地址才是shape类型的数据？
```
//p is 12.560000	,s is 12.560000
printInformation(&c1)
```
只要将c1的地址当参数传入printInformation函数中，系统不会报错！
因此必须明确接口的实现者到底是该种类型的数组值还是其地址！

### 模拟泛型
泛型是指 定义变量时，不明确该种变量的数据类型，必须等到调用时才能明确类型。
泛型在编程时是十分有用，但是Go中没有泛型类型，可以使用接口模拟泛型
函数可以接受接口类型的数据，那么如果接受一个空类型的数据会发生什么?
```
func printType(e interface{}){
	fmt.Printf("e type is %T\n",e)
}
```
printType接受空类型的参数，可以看看其调用：
```
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
```

printType可以接受整数类型、浮点数类型、字符串类型、数组类型数据，这意味着空接口数据是没有限制的，
**任何类型数据都是空接口类型数据！**
空接口，意味着其内部没有要求任何方法实现，也就是没有限制，也就意味着一切数据都是该抽象类型！

### 编写自定义服务器的API
接口的典型例子就是创建服务器 API.
编写 Web 服务器的常用方式是使用 net/http 程序包中的 http.Handler 接口
接口说明如下:
```go
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Handler) error
```
ListenAndServe 函数需要服务器地址（如 http://localhost:8000）以及将响应从调用调度至服务器地址的 Handler 的实例
编写自定义API:
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string{
	//Sprintf 返回格式内容的字符串 $float
	return fmt.Sprintf("$%.2f",d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter,r *http.Request){
	for item,price := range db{
		// Fprintf 将格式内容的字符串写入w
		fmt.Fprintf(w,"%s:%s\n",item,price)
	}
}

func main() {
	db := database{"Go T-Shift":25,"Go Jack":55}
	log.Fatal(http.ListenAndServe("localhost:8000",db))
}
```

# 在Go中使用并发
学习目标：
- Go中并发的工资原理
- 并发与并行之间的差异
- 如果在并发程序中使用channel（通道）进行通信
- 如何通过并发编写运行速度更快的程序
- 如果编写可以使用缓冲区来利用并发的动态程序

## 了解goroutine(轻量线程)
编写并发程序时最大的问题是在进程之间共享数据。 Go 采用不同于其他编程语言的通信方式，因为 Go 是通过 channel 来回传递数据的
不是通过共享内存通信，而是通过通信共享内存

### goroutine
goroutine 是轻量线程中的并发活动，而不是在操作系统中进行的传统活动
创建goroutine ，则必须在调用函数之前使用go关键字

不适用goroutine来处理调试网站：
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _,api := range apis{
		_,err := http.Get(api)
		if err != nil{
			fmt.Printf("Error:%s is down\n",err)
			continue
		}
		fmt.Printf("Success:%s is up and running!\n",api)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n",elapsed)
}
```
输出：
```
Success:https://management.azure.com is up and running!
Success:https://dev.azure.com is up and running!
Success:https://api.github.com is up and running!
Success:https://outlook.office.com/ is up and running!
Error:Get "https://api.somewhereintheinternet.com/": dial tcp: lookup api.somewhereintheinternet.com: no such host is down
Success:https://graph.microsoft.com is up and running!
Done! It took 8.380591s seconds
```

总花费时间接近8.5秒，将这个程序改成goroutine来处理：
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _,api := range apis{
		go checkApi(api)
	}
	// 由于处理是在太快了 ，加上休眠2s来看输出结果
	time.Sleep(time.Second * 2)
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n",elapsed)
}

// 由于go关键字必须加在函数前面，将一些逻辑提出来成为一个函数
func checkApi(api string){
	_,err:= http.Get(api)
	if err != nil{
		fmt.Printf("Error:%s is down\n",err)
		return
	}
	fmt.Printf("Success:%s is up and running!\n",api)
}
```
花费时间：
```
Error:Get "https://api.somewhereintheinternet.com/": dial tcp: lookup api.somewhereintheinternet.com: no such host is down
Success:https://dev.azure.com is up and running!
Success:https://management.azure.com is up and running!
Success:https://api.github.com is up and running!
Success:https://graph.microsoft.com is up and running!
Success:https://outlook.office.com/ is up and running!
Done! It took 2.0086604s seconds
```

可以看出来，即使休眠了2s，总时间也只花费了2.0086s,也就是说程序只花费了0.0086s
比之前不适用并发的程序速度提高了1000倍
但是由于程序处理过快，如果不加休眠是输出是没有结果的，因此为了方便管理goroutine，引起了channel

## 将channel 用作通信机制
Go 中的 channel 是 goroutine 之间的通信机制
当你需要将值从一个 goroutine 发送到另一个 goroutine 时，将使用 channel。

### channel语法
由于 channel 是发送和接收数据的通信机制，因此它也有类型之分，因此只能发送channel支持的数据类型
除使用关键字 chan 作为 channel 的数据类型外，还需指定将通过 channel 传递的数据类型，如 int 类型。
因此，chan int 便是 整数类型的通道，这两个一起用才能表示一种数据类型

使用make创建通道
```
ch := make(chan int)
```
一个channel 具有两个操作，发送数据和接受数据。若要指定channel具有的操作类型，需要使用运算发 <- 。
希望 channel 仅发送数据，则必须在 channel 之后使用 <- 运算符。 如果希望 channel 接收数据，则必须在 channel 之前使用 <- 运算符
```
ch <- x  //ch 发送通道，将x 通过通道ch 发送出去
x <- ch // ch 接受通道， 通过通道ch 接受数据x
```

关闭通道：
```
close(ch)
```

### 使用channel 修改网站调式程序
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	for _,api := range apis{
		go checkApi(api,ch)
	}
	// 输出 通道中的数据
	fmt.Print(<- ch)
	end := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n",end)
}

func checkApi(api string,ch chan string){
	_,err := http.Get(api)
	if err != nil{
		// 将格式化的诗句发送给ch通道
		ch <- fmt.Sprintf("ERROR:%s is down\n",api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS:%s is up and running!\n",api)
}
```
输出:
```
SUCCESS:https://management.azure.com is up and running!
Done! It took 597.5514ms seconds
```
不加休眠函数，也可以看到输出，但是只能看到其中一个 goroutine 的输出，而我们共创建了五个 goroutine。
这是因为make函数默认创建一个无缓冲区的channel,无缓冲区的channel会阻止发送行为，直到有人准备好接收数据
因此可以说发送和接受都是阻止操作，会中断程序执行
### 无缓冲区的channel
在程序中:
```
fmt.Print(<- ch)
```
我们可以说 fmt.Print(<-ch) 会阻止程序，因为它从 channel 读取，并等待一些数据到达。
一旦有任何数据到达，它就会继续下一行，然后程序完成。
其他 goroutine 发生了什么？ 它们仍在运行，但都没有在侦听。 而且，由于程序提前完成，一些 goroutine 无法发送数据。
因此可以说，发送通道和接受通道是一一对应关系，一个发送通道中ch发送数据就会等到 接收通道中ch接受数据，类似与通道的两端
这类似与ES7中的 async 和 await 关键字。

如果我们加一行 fmt.print(<- ch) :
```
// 输出 通道中的数据
fmt.Print(<- ch)
fmt.Print(<- ch)
```
会看到输出多了一行：
```
ERROR:https://api.somewhereintheinternet.com/ is down
SUCCESS:https://management.azure.com is up and running!
```
因此，为了正确看到输出，就必须确定goroutine的数量，修改后的程序为：
```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	for _,api := range apis{
		go checkApi(api,ch)
	}
	// 输出 通道中的数据
	// 修改后的部分，确定与发送通道相比配的接受通道
	for i:=0;i<len(apis);i++{
		fmt.Print(<- ch)
	}

	end := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n",end)
}

func checkApi(api string,ch chan string){
	_,err := http.Get(api)
	if err != nil{
		// 将格式化的诗句发送给ch通道
		ch <- fmt.Sprintf("ERROR:%s is down\n",api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS:%s is up and running!\n",api)
}
```

输出：
```
SUCCESS:https://management.azure.com is up and running!
SUCCESS:https://api.github.com is up and running!
SUCCESS:https://graph.microsoft.com is up and running!
SUCCESS:https://outlook.office.com/ is up and running!
ERROR:https://api.somewhereintheinternet.com/ is down
SUCCESS:https://dev.azure.com is up and running!
Done! It took 3.0608742s seconds
```
将输出结果与apis切片比对一下，可以发现，输出的内容顺序不是原来的顺序了，如果多次运行此程序会发现，基本上每次顺序都不一样。
这是因为无缓冲区的channel是同步操作的，它们保证每次发送数据时，程序都会被阻止，直到有人从 channel 中读取数据。

### 有缓冲区的channel
语法：
```
ch := make(chan string, 10)
```

### 指定channel方向
Go 中 channel 的一个有趣特性是，在使用 channel 作为函数的参数时，可以指定 channel 是要发送数据还是接收数据
语法：
```
chan<- int // 只能发送数据
<-chan int // 只能接收数据
```
比如：
```go
package main

import "fmt"

func send(ch chan<- string,message string ){
	fmt.Printf("Sending:%s\n",message)
	ch <- message
}

func read(ch <-chan string){
	fmt.Printf("Receiving:%v",<- ch)
}

func main() {
	ch := make(chan string,1)
	// 不能使用goroutine 因为send和read是有顺序的，使用go 关键字后就会出现顺序错误
	send(ch,"Hello world")
	read(ch)
}
```