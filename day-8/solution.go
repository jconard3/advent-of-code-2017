package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const inputName = "test.txt"

func main() {
	//open input
	f, err := os.Open(inputName)
	if err != nil {
		fmt.Printf("Could not open file %s. Exiting.", inputName)
		os.Exit(-1)
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		r := regexp.MustCompile(`(?P<Register>\w+) (?P<RegMod>\w+) (?P<Amount>(\-)?\d+) if `)
	}
}
