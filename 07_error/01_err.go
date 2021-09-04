package main

import (
	"errors"
	"fmt"
	"time"
)

type employee struct {
	id int
	firstname string
	lastname string
	address string
}

var errnotFound = errors.New("employee not found")

func getInformationById(id int) (*employee,error){
	if id != 1000{
		return nil,errnotFound
	}
	emp := employee{lastname: "Doe",firstname: "Jhon"}
	return &emp,nil
}

func getInformation(id int) (*employee,error){
	for tries := 0;tries<3;tries++{
		emp,err := apiCallEmployee(1000)
		if err == nil{
			return emp,nil
		}
		fmt.Println("Server is not responding,retrying...")
		time.Sleep(time.Second * 2)
	}
	return nil,fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*employee,error){
	e := employee{firstname: "Doe",lastname:"Jhon"}
	return &e,nil
}