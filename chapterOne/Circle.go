package main

import "math"

type Circle struct {
	x, y, r float64
}

func (d *Circle) CircleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
