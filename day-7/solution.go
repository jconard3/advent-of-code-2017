package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Program struct {
	name     string
	weight   int
	children []string
}

const inputFile = "input.txt"

func main() {
	//open our input
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Cannot open file %s. Exiting\n", inputFile)
		os.Exit(-1)
	}

	//initialize the data structure holding file contents
	var programs = []Program{}
	var allChildren = []string{} //collection of all children
	//start processing our file
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()     //grab the current line
		var p = Program{} // program struct to hold values from line
		r := regexp.MustCompile(`(?P<Name>\w+) \((?P<Weight>\d+)\)(?: -> (?P<Children>.*))?`)
		subs := r.FindStringSubmatch(l) //[1] = Name, [2] = Weight, [3] = Children
		p.name = subs[1]
		p.weight, _ = strconv.Atoi(subs[2])
		//Processing of the children
		if subs[3] != "" { //if the regexp matches children, add them to the struct
			p.children = strings.Split(subs[3], ",")
			//Trim space on all children in struct
			for i := range p.children {
				p.children[i] = strings.TrimSpace(p.children[i])
			}
		}
		programs = append(programs, p) //save our temporary program in to the larger slice
		//copy the children from this line into larger collection
		if len(p.children) > 0 {
			allChildren = appendChildren(allChildren, p.children)
		}
	}

	//Algorithm: For each program in 'programs'
	// retrieve the program 'name'
	// if program 'name' exists in 'allChildren'
	// then program is not root program
	for i := range programs {
		n := programs[i].name
		for j := 0; j <= len(allChildren); j++ {
			if j == len(allChildren) { //if we get to the end of allChildren, program name is not a child and therefore the root
				fmt.Printf("The root program name is %s", n)
				os.Exit(0)
			}
			if n == allChildren[j] { //if the name of the program is a child, it is not the root
				break
			}
		}
	}
	fmt.Println("Root program name not found.")
	os.Exit(-1)
}

//function appropriated from https://blog.golang.org/go-slices-usage-and-internals
func appendChildren(allChildren []string, newChildren []string) []string {
	m := len(allChildren)     //previous length
	n := m + len(newChildren) // new needed length
	if n > cap(allChildren) { // if new needed length exceeds existing capacity
		newSlice := make([]string, (n + 1)) // +1 in case n = 0
		copy(newSlice, allChildren)
		allChildren = newSlice
	}
	allChildren = allChildren[0:n]      // expand slice length to meet new needed length
	copy(allChildren[m:n], newChildren) // copy the new Children to the 'top' of allChildren
	return allChildren
}
