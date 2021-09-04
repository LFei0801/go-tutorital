package main

import "fmt"

type shape interface {
	perimeter() float64
	area() float64
}

type circle struct {
	radius float64
}

func (c *circle) perimeter() float64{
	return 2 * 3.14 * c.radius
}

func (c *circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func printInformation(s shape){
	fmt.Printf("%T\n",s)
	fmt.Printf("p is %f\t,s is %f\n",s.perimeter(),s.area())
}

func printType(e interface{}){
	fmt.Printf("e type is %T\n",e)
}


