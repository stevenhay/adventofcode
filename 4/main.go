package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.OpenFile("input.txt", os.O_RDONLY, 0755)
	defer input.Close()

	scanner := bufio.NewScanner(input)

	part1 := 0
	part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		ls, le, rs, re := GetRange(line)

		if (ls >= rs && le <= re) || (rs >= ls && re <= le) {
			part1++
		}

		if (ls >= rs && ls <= re) || (rs >= ls && rs <= le) {
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func GetRange(line string) (int, int, int, int) {
	e := strings.Split(line, ",")

	e1 := strings.Split(e[0], "-")
	ls, _ := strconv.Atoi(e1[0])
	le, _ := strconv.Atoi(e1[1])

	e2 := strings.Split(e[1], "-")
	rs, _ := strconv.Atoi(e2[0])
	re, _ := strconv.Atoi(e2[1])

	return ls, le, rs, re
}
