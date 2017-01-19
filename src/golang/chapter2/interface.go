package main

import (
	"fmt"
	"math"
)

func main() {
	r := Rectangle{width: 3, height: 4}
	c := Circle{radius: 2}
	measure(r)
	measure(c)
}

type IGeometry interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (cir Circle) area() float64 {
	return cir.radius * cir.radius * math.Pi
}

type Rectangle struct {
	width, height float64
}

func (rec Rectangle) area() float64 {
	return rec.width * rec.height
}

func measure(geo IGeometry) {
	fmt.Println(geo)
	fmt.Println(geo.area())
}
