package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	defer input.Close()

	scanner := bufio.NewScanner(input)
	part2(scanner)
}

func part2(scanner *bufio.Scanner) {
	prioSum := 0
	index := 0
	var strlst [3]string

	for scanner.Scan() {
		line := scanner.Text()
		strlst[index] = line
		index++

		if index == 3 {
			prioSum += GetPriority(FindBadge(strlst))
			index = 0
		}
	}
	fmt.Printf("Total Prio: %d", prioSum)
}

func FindBadge(str [3]string) rune {
	m := make(map[rune]int)
	for _, s := range str {
		lm := make(map[rune]bool)
		for _, c := range s {
			if _, ok := lm[c]; !ok {
				lm[c] = true

				if i, ok := m[c]; ok {
					if i+1 == 3 {
						return c
					}
					m[c] = i + 1
				} else {
					m[c] = 1
				}
			}
		}

	}
	return -1
}

func part1(scanner *bufio.Scanner) {
	prioSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		h1 := line[:len(line)/2]
		h2 := line[len(line)/2:]
		fmt.Printf("%s : %s\n", h1, h2)

		m := make(map[rune]bool)
		for _, c := range h1 {
			m[c] = false
		}
		for _, c := range h2 {
			if v, ok := m[c]; ok && !v {
				m[c] = true
				prioSum += GetPriority(c)
			}
		}
	}

	fmt.Printf("Total Prio: %d", prioSum)
}

func GetPriority(x rune) int {
	ascii := int(x)
	if ascii >= 65 && ascii <= 90 {
		return 27 + 1*(ascii-65)
	}
	return 1 + 1*(ascii-97)
}
