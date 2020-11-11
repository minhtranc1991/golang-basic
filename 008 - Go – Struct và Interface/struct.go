package main

import (
	"fmt"
	"math"
)

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	r := Rectangle{x1: 0, y1: 10, x2: 0, y2: 10}
	fmt.Println(circleArea(&c))
	fmt.Println(rectangleArea(&r))
	fmt.Println(c.area())
	fmt.Println(r.area())
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func rectangleArea(r *Rectangle) float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
