package main

type Point struct {
	x, y int
}

func (p Point) northOf() Point {
	return Point{p.x, p.y - 1}
}

func (p Point) eastOf() Point {
	return Point{p.x + 1, p.y}
}

func (p Point) southOf() Point {
	return Point{p.x, p.y + 1}
}

func (p Point) westOf() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) oneBased() Point {
	return Point{p.x + 1, p.y + 1}
}

func (p Point) getAdjacent(d Direction) Point {
	var spot Point
	switch d {
	case North:
		spot = p.northOf()
	case East:
		spot = p.eastOf()
	case South:
		spot = p.southOf()
	case West:
		spot = p.westOf()
	default:
		panic("Diagonals not supported")
	}
	return spot
}
