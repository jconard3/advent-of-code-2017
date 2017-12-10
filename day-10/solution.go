package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileInput = "input.txt"

const listLen = 256

//const listLen = 5 //list length for test.txt solution

func main() {
	//initialize resources
	//list initialization
	list := ring.New(listLen)
	for i := 0; i < list.Len(); i++ {
		list.Value = i
		list = list.Next()
	}

	totalSkip := 0 //tracker for how much we've skipped down the ring. This will help us navigate back to the 'beginning' of the ring for the final answer to part 1
	skipSize := 0  //skip size initializaiton

	//lengths initialization
	f, err := os.Open(fileInput)
	if err != nil {
		fmt.Printf("Could not open file %s", fileInput)
		os.Exit(-1)
	}
	s := bufio.NewScanner(f)
	s.Scan() //input should only be one line
	l := s.Text()
	ls := strings.Split(l, ",")
	lengths := make([]int, len(ls))
	for i := range ls {
		lengths[i], _ = strconv.Atoi(strings.TrimSpace(ls[i]))
	}

	//hash computation
	for i := range lengths {
		length := lengths[i]
		subList := make([]int, length)
		//populate subList and move ring pointer to the end of the twist
		for j := 0; j < length; j++ {
			subList[j] = list.Value.(int)
			list = list.Next()
		}
		//reverse down the ring and populate from the front of the subList
		for j := 0; j < length; j++ {
			list = list.Prev()
			list.Value = subList[j]
		}
		//Move position on ring forward by the length + skipSize
		forward := length + skipSize
		totalSkip = totalSkip + forward
		for i := 0; i < forward; i++ {
			list = list.Next()
		}
		//Increase the skipSize by one
		skipSize = skipSize + 1
	}

	//reverse the ring back the totalSkip size to get back to the 'first' element
	for i := 0; i < totalSkip; i++ {
		list = list.Prev()
	}
	solution := list.Value.(int)
	fmt.Printf("First ring value = %d\n", list.Value.(int))
	list = list.Next()
	solution = solution * list.Value.(int)
	fmt.Printf("Second ring value = %d\n", list.Value.(int))
	fmt.Printf("Solution = %d\n", solution)
}
