package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fileInput = "test.txt"

func main() {
	line := make(map[string]int)
	line["a"] = 0
	line["b"] = 1
	line["c"] = 2
	line["d"] = 3
	line["e"] = 4

	//open up input
	f, err := os.Open(fileInput)
	if err != nil {
		fmt.Printf("Could not open file %s. Exiting\n", fileInput)
		os.Exit(-1)
	}
	s := bufio.NewScanner(f)
	s.Scan()      //test is only on first line
	l := s.Text() //get text of first line
	commands := strings.Split(l, ",")
	for _, command := range commands {
		switch move := string(command[0]); move {
		case "s":
			size, _ := strconv.Atoi(string(command[1]))
			spin(line, size)
		case "x":
			sub := strings.Split(command[1:], "/")
			first, _ := strconv.Atoi(sub[0])
			second, _ := strconv.Atoi(sub[1])
			exchange(line, first, second)
		case "p":
			sub := strings.Split(command[1:], "/")
			first := sub[0]
			second := sub[1]
			partner(line, first, second)
		}
	}
	fmt.Println(line)
}

func spin(line map[string]int, size int) {
	for i := size; i > 0; i-- {
		for prog, pos := range line {
			line[prog] = (pos + 1) % (len(line))
		}
	}
}

func exchange(line map[string]int, first int, second int) {
	for prog, pos := range line {
		if pos == first {
			line[prog] = second
		} else if pos == second {
			line[prog] = first
		}
	}
}

func partner(line map[string]int, first string, second string) {
	firstPos := line[first]
	secondPos := line[second]
	//switch first and second positions
	line[first] = secondPos
	line[second] = firstPos
}
