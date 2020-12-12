package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type coord struct {
	X int
	Y int
}

type position struct {
	location  coord
	direction int
}

func main() {
	scanner := bufio.NewScanner(openStdinOrFile())

	shipPos := position{
		coord{0, 0},
		90,
	}

	waypoint := coord{10, 1}

	for scanner.Scan() {
		instruction := scanner.Text()

		direction := instruction[0]
		magnitude, err := strconv.Atoi(instruction[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch direction {
		case 'N':
			waypoint.Y += magnitude
		case 'E':
			waypoint.X += magnitude
		case 'S':
			waypoint.Y -= magnitude
		case 'W':
			waypoint.X -= magnitude
		case 'R':
			waypoint.rotate(-magnitude)
		case 'L':
			waypoint.rotate(magnitude)
		case 'F':
			shipPos.moveAbs(waypoint.X*magnitude, waypoint.Y*magnitude)
		}
		fmt.Printf("New Ship Position: %v, waypoint: %v\n", shipPos, waypoint)
	}
	fmt.Printf("The hamilton distance is: %v", math.Abs(float64(shipPos.location.X))+math.Abs(float64(shipPos.location.Y)))
}

func (p *position) moveAbs(x int, y int) {
	p.location.X += x
	p.location.Y += y
}

func (p *position) moveRel(delta int, theta int) {
	p.direction = ((p.direction + theta) + 360) % 360
	if p.direction == 0 {
		p.location.X += delta
	} else if p.direction == 90 {
		p.location.Y += delta
	} else if p.direction == 180 {
		p.location.X -= delta
	} else if p.direction == 270 {
		p.location.Y -= delta
	}
}

func (c *coord) rotate(degrees int) {
	var sin, cos int
	if (degrees+360)%360 == 0 {
		sin = 0
		cos = 1
	} else if (degrees+360)%360 == 90 {
		sin = 1
		cos = 0
	} else if (degrees+360)%360 == 180 {
		sin = 0
		cos = -1
	} else if (degrees+360)%360 == 270 {
		sin = -1
		cos = 0
	}

	x := c.X*cos - c.Y*sin
	y := c.X*sin + c.Y*cos

	c.X = x
	c.Y = y
}

func openStdinOrFile() io.Reader {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return r
}
