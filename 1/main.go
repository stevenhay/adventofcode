package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, _ := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	calories := make([]int, 0)
	currCals := 0
	for scanner.Scan() {
		t := scanner.Text()
		// fmt.Println(t)
		if t == "" {
			fmt.Printf("Elf = %d\n", currCals)
			calories = append(calories, currCals)
			currCals = 0
		}

		c, _ := strconv.Atoi(t)
		currCals += c
	}

	sort.Ints(calories)
	fmt.Printf("Max: %d\n", calories[len(calories)-1])
	fmt.Printf("Max-1: %d\n", calories[len(calories)-2])
	fmt.Printf("Max-2: %d\n", calories[len(calories)-3])
}
