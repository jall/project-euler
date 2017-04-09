package main

import "fmt"

type Coord struct {
	x int
	y int
}

type Grid map[Coord]int

func main() {
	maxRadius := 500
	spiral := createSpiral(maxRadius)
	sum := sumOfDiagonals(spiral, maxRadius)
	fmt.Println(sum)
}

func createSpiral(maxRadius int) Grid {
	spiral := Grid{}

	n := 1
	x, y := 0, 0

	spiral[Coord{x, y}] = n

	for radius := 1; radius <= maxRadius; radius++ {
		// Move 1 to Right for new radius
		x++
		n++
		spiral[Coord{x, y}] = n

		// Down
		for {
			y--
			n++
			spiral[Coord{x, y}] = n

			if -y >= radius {
				break
			}
		}

		// Left
		for {
			x--
			n++
			spiral[Coord{x, y}] = n

			if -x >= radius {
				break
			}
		}

		// Up
		for {
			y++
			n++
			spiral[Coord{x, y}] = n

			if y >= radius {
				break
			}
		}

		// Right
		for {
			x++
			n++
			spiral[Coord{x, y}] = n

			if x >= radius {
				break
			}
		}
	}

	return spiral
}

func sumOfDiagonals(spiral Grid, radius int) int {
	sum := spiral[Coord{0, 0}]

	for i := 1; i <= radius; i++ {
		sum += spiral[Coord{i, i}]
		sum += spiral[Coord{i, -i}]
		sum += spiral[Coord{-i, i}]
		sum += spiral[Coord{-i, -i}]
	}

	return sum
}
