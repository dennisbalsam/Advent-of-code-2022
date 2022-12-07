package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(data)
	startOfPacketMarker := 0
	startOfMessageMarker := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := []string{string(line[0]), string(line[1]), string(line[2])}

		for i, v := range line[3:] {
			if ok, index := containsAtPosition(values, string(v)); ok {
				values = append(values[index+1:], string(v))
			} else {
				values = append(values, string(v))
				if len(values) == 4 && startOfPacketMarker == 0 {
					startOfPacketMarker = i + 4
				}
				if len(values) == 14 && startOfMessageMarker == 0 {
					startOfMessageMarker = i + 4
				}
			}
		}
	}
	// Part 1
	fmt.Printf("Total number of characters needed for packet marker: %v\n", startOfPacketMarker)

	// Part 2
	fmt.Printf("Total number of characters needed for message marker: %v", startOfMessageMarker)
}

func containsAtPosition(s []string, v string) (bool, int) {
	for i, a := range s {
		if a == v {
			return true, i
		}
	}
	return false, 0
}
