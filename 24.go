package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x, y float64
}

func (p1 *Point) distance(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2.0) + math.Pow(p2.y-p1.y, 2.0))
}

func New(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {

	point1 := New(4, 500)
	point2 := New(5, 10)

	fmt.Printf("point1 | x %f; y %f\n", point1.x, point1.y)
	fmt.Printf("point2 | x %f; y %f\n", point2.x, point2.y)

	fmt.Printf("Расстояние между точками = %f", point1.distance(point2))
}
