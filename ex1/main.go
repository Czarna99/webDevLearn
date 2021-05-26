package main

import "fmt"

const pi = 3.14

type square struct {
	side int
}
type circle struct {
	radius int
}

func (s square) area() float64 {
	return float64(s.side * s.side)

}

func (c circle) area() float64 {
	return float64(c.radius*c.radius) * pi
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{15}
	cir := circle{7}

	info(sq)
	info(cir)

}
