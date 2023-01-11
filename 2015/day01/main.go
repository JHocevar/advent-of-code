package main

import (
	_ "embed"
	"fmt"

	"github.com/jhocevar/advent-of-code/util"
)

//go:embed input.txt
var input string

func main() {
	ans := util.RunPart(part1, part2, input)
	fmt.Println(ans)
}

func part1(input string) int {
	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
	}
	return floor
}

func part2(input string) int {
	floor := 0
	for count, c := range input {
		if c == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			return count + 1
		}
	}
	return -1
}
