package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	values := getPriorityMapping()
	data, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// part 1 values
	sum := 0

	//part 2 values
	groupsOf3Sum, groupDivider := 0, 1
	groupValues := map[string]map[string]string{"firstGroupValues": map[string]string{}, "mutualValues": map[string]string{}}

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lineValues := map[string]string{}

		line := scanner.Text()
		for i, c := range line {
			// part 1 logic
			if i+1 <= len(line)/2 {
				lineValues[string(c)] = ""
			} else {
				if _, ok := lineValues[string(c)]; ok {
					sum += values[string(c)]
					delete(lineValues, string(c))
				}
			}

			// part 2 logic
			switch groupDivider {
			case 1:
				groupValues["firstGroupValues"][string(c)] = ""
			case 2:
				if _, ok := groupValues["firstGroupValues"][string(c)]; ok {
					groupValues["mutualValues"][string(c)] = ""
					delete(groupValues["firstGroupValues"], string(c))
				}
			case 3:
				if _, ok := groupValues["mutualValues"][string(c)]; ok {
					groupsOf3Sum += values[string(c)]
					delete(groupValues["mutualValues"], string(c))
				}
			}
		}
		if groupDivider == 3 {
			groupDivider = 0
			groupValues = map[string]map[string]string{"firstGroupValues": map[string]string{}, "mutualValues": map[string]string{}}
		}
		groupDivider += 1
	}

	fmt.Printf("Total Sum Part 1: %v\n", sum)
	fmt.Printf("Total Sum Part 2: %v\n", groupsOf3Sum)

}

func getPriorityMapping() map[string]int {
	values := map[string]int{}
	priority := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		values[fmt.Sprintf("%c", ch)] = priority
		priority += 1
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		values[fmt.Sprintf("%c", ch)] = priority
		priority += 1
	}
	return values
}
