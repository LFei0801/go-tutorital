package main

import (
	"fmt"
	"net/http"
	"time"
)

func goroutineBasis() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		go checkApiInGoroutine(api)
	}
	time.Sleep(time.Second * 2)
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds\n", elapsed)
}

func checkApiInGoroutine(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("Error:%s is down\n", err)
		return
	}
	fmt.Printf("Success:%s is up and running!\n", api)
}
