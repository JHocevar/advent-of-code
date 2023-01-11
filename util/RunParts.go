package util

import (
	"flag"
	"log"
)

func RunPart[I any](part1 func(I) int, part2 func(I) int, input I) int {
	part := flag.Int("part", 1, "Which part of the question to run")
	flag.Parse()

	var ans int
	if *part == 1 {
		ans = part1(input)
	} else if *part == 2 {
		ans = part2(input)
	} else {
		log.Fatal("Part must be 1 or 2")
	}

	return ans
}
