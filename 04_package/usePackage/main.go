package main

import "../calculator"

func main() {
	total := calculator.Sum(3,5)
	println(total) //8
	println("Version:",calculator.Version) //Version: 1.0
}

