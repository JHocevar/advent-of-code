package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
)

func main() {
	year := flag.Int("year", 2015, "Year of Advent of Code to run")
	day := flag.Int("day", 1, "Day of Advent of Code to run")
	part := flag.Int("part", 1, "Part of Advent of Code question to run")
	flag.Parse()

	var dayString string
	if *day < 10 {
		dayString = "0" + strconv.Itoa(*day)
	} else {
		dayString = strconv.Itoa(*day)
	}

	filePath := "./" + strconv.Itoa(*year) + "/day" + dayString + "/main.go"

	command := exec.Command("go", "run", filePath, "-part", strconv.Itoa(*part))
	output, err := command.Output()
	if err == nil {
		fmt.Print(string(output))
	} else {
		fmt.Println("an error occured:", err)
	}
}
