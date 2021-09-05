package main

import (
	"fmt"
	"time"
)

func fib(number int,ch chan<- string){
	x,y := 1,1
	for i:=0;i<=number;i++{
		x,y = y,x+y
	}
	ch <- fmt.Sprintf("Fib(%v) is %v\n",number,x)
}

func fibNoGo (number int)  int{
	x,y := 1,1
	for i:=0;i<=number;i++{
		x,y = y,x+y
	}
	return x
}

func main() {
	start := time.Now()

	//for i:=0;i<1000;i++{
	//	fmt.Printf("Fib(%v) is %v\n",i,fibNoGo(i))
	//}
	ch := make(chan string,1000)

	for i:=0;i<1000;i++{
		go fib(i,ch)
	}

	for i:=0;i<1000;i++{
		fmt.Printf(<- ch)
	}

	end := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n",end)
}