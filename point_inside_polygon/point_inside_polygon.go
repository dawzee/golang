package main

import (
	"fmt"
)

type vertex struct {
	x, y int
}

func sub (a, b vertex) vertex {
	return vertex{a.x-b.x, a.y-b.y}
}

func product (a, b vertex) int {
	return (a.x*b.y)-(a.y*b.x)
}

func main() {
	poly := make([]vertex,4)
	poly[0] = vertex{0, 0}
	poly[1] = vertex{4, 0}
	poly[2] = vertex{4, 4}
	poly[3] = vertex{0, 4}


	fmt.Println("0,0 = ", inside_convex_polygon(vertex{0,0},poly))
	fmt.Println("1,1 = ", inside_convex_polygon(vertex{1,1},poly))
	fmt.Println("1,2 = ", inside_convex_polygon(vertex{1,2},poly))
	fmt.Println("1,3 = ", inside_convex_polygon(vertex{1,3},poly))
	fmt.Println("1,4 = ", inside_convex_polygon(vertex{1,4},poly))
}

func inside_convex_polygon(point vertex, poly []vertex) bool {

	previous_side := 0

	for i := range poly {
		a := poly[i]
		b := poly[(i+1)%len(poly)]
		current_side := product(sub(b, a), sub(point, a))
		if current_side == 0 {
			return false
		} else if previous_side == 0 {
			previous_side = current_side
		} else if (current_side > 0 && previous_side < 0) ||
			(current_side < 0 && previous_side > 0) {
			return false
		}
	}
	return true
}
