package main

import (
	"fmt"
	"log"
	"os"
)

func logBasis(){
	log.Print("Hey,i'm a log")
}

func logFatal(){
	log.Fatal("Hey,i'm a log by fatal")
	fmt.Print("Can you see me!")
}

func logPanic(){
	log.Panic("Hey,i'm a log error by panic")
	fmt.Println("Can you see me")
}

func logSetPrefix(){
	log.SetPrefix("添加的前缀():")
	log.Print("i'm a log ")
	log.Fatal("i'm a fatal log")
	fmt.Println("Can you see me")
}

func logToFile(){
	// 创建日志文件，并配置权限
	file,err := os.OpenFile("info.log",os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	// 配置记录到文件
	log.SetOutput(file)

	log.Print("Hey,i'm a log")
}