package main

type Direction int

const (
	North Direction = iota
	Northeast
	East
	Southeast
	South
	Southwest
	West
	Northwest
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case Northeast:
		return "Northeast"
	case East:
		return "East"
	case Southeast:
		return "Southeast"
	case South:
		return "South"
	case Southwest:
		return "Southwest"
	case West:
		return "West"
	case Northwest:
		return "Northwest"
	default:
		return "Unknown"
	}
}

func (d Direction) turnRight90() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	case Northeast, Southeast, Southwest, Northwest:
		panic("Diagonals not supported")
	default:
		panic("Diagonals not supported")
	}
}
