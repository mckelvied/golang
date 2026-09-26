package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base, height float64
}

type square struct {
	side float64
}

func (t triangle) getArea() float64 {
	if t.base <= 0 || t.height <= 0 {
		return 0
	}
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	if s.side <= 0 {
		return 0
	}
	return s.side * s.side
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	s1 := square{side: 5}
	t1 := triangle{base: 3, height: 4}
	printArea(s1)
	printArea(t1)

}
