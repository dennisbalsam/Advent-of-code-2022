package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fullyOverlapSum := 0
	anyOverlapSum := 0
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()

		ranges := map[int][]int{}
		for i, seq := range strings.Split(line, ",") {
			for _, num := range strings.Split(seq, "-") {
				n, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal("Could not convert string")
				}
				ranges[i+1] = append(ranges[i+1], n)
			}
		}

		// part 1
		if ranges[1][0] <= ranges[2][0] && ranges[1][1] >= ranges[2][1] {
			fullyOverlapSum += 1
		} else if ranges[2][0] <= ranges[1][0] && ranges[2][1] >= ranges[1][1] {
			fullyOverlapSum += 1
		} else if (ranges[1][1] >= ranges[2][0] && ranges[1][1] <= ranges[2][1]) || (ranges[2][1] >= ranges[1][0] && ranges[2][1] <= ranges[1][1]) { // part 2
			anyOverlapSum += 1
		}

	}

	// part 1
	fmt.Printf("Total Sum of Fully Overlapped pairs: %v\n", fullyOverlapSum)

	// part 1
	fmt.Printf("Total Sum of Any Overlapped pairs: %v\n", anyOverlapSum+fullyOverlapSum)

}
