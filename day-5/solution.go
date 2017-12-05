package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Read input, convert to ints and place into array
	f, err := os.Open("input.txt")
	//f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Could not open input.txt. Exiting")
		os.Exit(-1)
	}
	defer f.Close()

	var list []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		i, _ := strconv.Atoi(s.Text())
		list = append(list, i)
	}

	//Start processing
	pos := 0   //Initialize the current position
	count := 0 //initialize the count
	length := len(list)
	for !(pos >= length) {
		count = count + 1
		jump := list[pos]
		if jump >= 3 {
			list[pos] = list[pos] - 1
		} else {
			list[pos] = list[pos] + 1
		}
		pos = pos + jump
	}
	fmt.Printf("It took %d steps", count)
}
