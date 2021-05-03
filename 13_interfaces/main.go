package main

import (
	"fmt"
	"math"
)

// Define interfact

type Shape interface {
	area() float64
}

type Circle struct {
	x,y, radius float64
}

type Rectangle struct {
	h, w float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.h * r.w 
}

func getArea(s Shape) float64 {
	return s.area()
}

func main (){ 
	circle := Circle{x:10, y:10, radius: 20}
	rectangle := Rectangle{h:20, w: 20}

	fmt.Printf("Circle Area %f\n", getArea(circle))
	fmt.Printf("Rectangle Area %f\n", getArea(rectangle))
}