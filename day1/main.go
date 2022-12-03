package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	max := 0                // for part 1
	top3 := [3]int{0, 0, 0} // for part 2
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(data)
	sum := 0
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			if sum > max {
				max = sum
			}
			for i, t := range top3 {
				if sum > t {
					top3[i] = sum
					sum = t
				}
			}
			sum = 0
			continue
		}
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Could not convert string to int")
		}

		sum += i
	}

	// Part 1 solution
	fmt.Printf("Highest number of calories for one group is: %v", max)

	//Part 2 solution
	sum = 0
	for _, v := range top3 {
		sum += v
	}
	fmt.Printf("\nNumber of calories for top 3 groups carrying most calories : %v", sum)
}
