package main

import (
	"fmt"
	"net/http"
	"time"
)

func channelBasis() {
	start := time.Now()
	// 无缓冲区 channel
	//ch := make(chan string)

	// 有缓冲区channel
	ch := make(chan string, 10)

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}
	for _, api := range apis {
		go checkApiByChannel(api, ch)
	}
	// 输出 通道中的数据
	// 修改后的部分，确定与发送通道相比配的接受通道
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	end := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n", end)
}

func checkApiByChannel(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		// 将格式化的诗句发送给ch通道
		ch <- fmt.Sprintf("ERROR:%s is down\n", api)
		return
	}
	ch <- fmt.Sprintf("SUCCESS:%s is up and running!\n", api)
}
