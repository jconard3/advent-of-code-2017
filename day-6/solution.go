package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Bank [16]int

func main() {
	//open file
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Could not open input file. Exiting")
		os.Exit(-1)
	}
	defer f.Close()

	//populate array representation of the banks
	var banks = Bank{}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords) //set scanner to split at space deliniated words
	for i := 0; i < len(banks); i++ {
		s.Scan()
		v, _ := strconv.Atoi(s.Text())
		banks[i] = v
	}

	//Start processing
	count := 0
	var record = []Bank{}
	record = append(record, banks)
	for {
		banks = redistribute(banks)
		count = count + 1
		if seenBefore(banks, record) {
			fmt.Printf("It took %d redistributions", count)
			break
		}
		record = append(record, banks)
	}
}

func redistribute(banks Bank) Bank {
	maxValue, maxIndex := maxSlice(banks) //Find max value and index
	banks[maxIndex] = 0                   //Reset value on max
	for i := 1; i <= maxValue; i++ {
		currentIndex := (maxIndex + i) % len(banks)
		banks[currentIndex] = banks[currentIndex] + 1
	}
	return banks
}

func maxSlice(banks Bank) (int, int) {
	var maxValue, maxIndex = 0, 0
	for i, v := range banks {
		if v > maxValue {
			maxValue = v
			maxIndex = i
		}
	}
	return maxValue, maxIndex
}

func seenBefore(banks Bank, record []Bank) bool {
	for i := range record {
		for j := range banks {
			if banks[j] != record[i][j] {
				break
			} else if j == len(banks)-1 {
				return true
			}
		}
	}
	return false
}
