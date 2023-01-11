package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jhocevar/advent-of-code/util"
)

//go:embed input.txt
var input string

func main() {
	ans := util.RunPart(part1, part2, parseInput(input))
	fmt.Println(ans)
}

func parseInput(input string) [][3]int {
	var outputLines [][3]int
	for _, val := range strings.Split(input, "\n") {
		outputLine := [3]int{0,0,0}
		for i, number := range strings.Split(val, "x") {
			// Not doing error catching here
			intNum, _ := strconv.Atoi(number)
			outputLine[i] = intNum
		}
		outputLines = append(outputLines, outputLine)
	}
	return outputLines
}

func part1(input [][3]int) int {
	total := 0
	for _, val := range input {
		l, w, h := val[0], val[1], val[2]
		total += 2*l*w + 2*w*h + 2*h*l

		sides := []int{l*w, w*h, h*l}
		sort.Ints(sides)
		total += sides[0]
	}
	return total
}

func part2(input [][3]int) int {
	total := 0
	for _, val := range input {
		l, w, h := val[0], val[1], val[2]
		total += l*w*h

		sort.Ints(val[:])
		total += 2*val[0] + 2*val[1]
	}
	return total
}
