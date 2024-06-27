package main

import "fmt"

type shape interface {
	getArea() float64
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

type square struct {
	sideLength float64
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

type triangle struct {
	height float64
	base   float64
}

func (tr triangle) getArea() float64 {
	return (tr.base * tr.height) * 0.5
}

func main() {

	printArea(square{10})
	printArea(square{5})
	printArea(triangle{5, 10})
	printArea(triangle{2, 3})

}
