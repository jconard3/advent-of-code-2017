package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "input.txt"

type Coordinate struct {
	X  float64
	Y  float64
	N  int
	NE int
	SE int
	S  int
	SW int
	NW int
}

func main() {
	//process input
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Could not open file %s. Exiting", filename)
		os.Exit(-1)
	}

	//scan file for a line of text
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		//process line of text
		coord := walk(l)
		fmt.Printf("The displacement walked from the origin is %g steps\n", coord.displacement())
	}
}

func walk(directions string) Coordinate {
	var coord = Coordinate{}
	dirs := strings.Split(directions, ",")
	for i := range dirs {
		switch d := dirs[i]; d {
		case "n":
			coord.N = coord.N + 1
		case "ne":
			coord.NE = coord.NE + 1
		case "se":
			coord.SE = coord.SE + 1
		case "s":
			coord.S = coord.S + 1
		case "sw":
			coord.SW = coord.SW + 1
		case "nw":
			coord.NW = coord.NW + 1
		default:
			fmt.Printf("WARN: Could not match direction %s in switch statement.", d)
		}
	}
	return coord
}

func (coord Coordinate) displacement() float64 {
	if coord.N > coord.S {
		coord.N = coord.N - coord.S
		coord.S = 0
	} else {
		coord.S = coord.S - coord.N
		coord.N = 0
	}
	if coord.NE > coord.SW {
		coord.NE = coord.NE - coord.SW
		coord.SW = 0
	} else {
		coord.SW = coord.SW - coord.NE
		coord.NE = 0
	}
	if coord.NW > coord.SE {
		coord.NW = coord.NW - coord.SE
		coord.SE = 0
	} else {
		coord.SE = coord.SE - coord.NW
		coord.NW = 0
	}

	//
	return 1
}
