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
	x, y := 0, 0
	m := make(map[string]int)
	m["0-0"] = 1
	for _, v := range input {
		switch v {
		case '<':
			x--
		case '^':
			y++
		case '>':
			x++
		case 'v':
			y--
		}
		m[fmt.Sprintf("%d-%d", x, y)] += 1
	}

	return len(m)
}

func part2(input string) int {
	cords := [2][2]int{}
	m := make(map[string]int)

	for i, v := range input {
		santa := i % 2
		switch v {
		case '<':
			cords[santa][0]--
		case '^':
			cords[santa][1]++
		case '>':
			cords[santa][0]++
		case 'v':
			cords[santa][1]--
		}
		coordString := fmt.Sprintf("%d-%d", cords[santa][0], cords[santa][1])
		m[coordString] += 1
	}

	return len(m)
}
