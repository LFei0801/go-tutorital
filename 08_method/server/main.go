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
