package day02

import (
	"github.com/tmslammi/go-aoc/util"
	"strings"
	"time"
)

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 2 part 1")
	fileContent := util.ReadFile("./input.txt")
	hPos := 0
	vPos := 0
	for i := 0; i < len(fileContent); i++ {
		parsed := strings.Split(fileContent[i], " ")
		dir := parsed[0]
		moveCount := util.ToInt(parsed[1])
		switch dir {
		case "down":
			hPos += moveCount
		case "up":
			hPos -= moveCount
		case "forward":
			vPos += moveCount
		}
	}
	return hPos * vPos
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 2 part 2")
	fileContent := util.ReadFile("./input.txt")
	aim := 0
	hPos := 0
	depth := 0
	for i := 0; i < len(fileContent); i++ {
		parsed := strings.Split(fileContent[i], " ")
		dir := parsed[0]
		moveCount := util.ToInt(parsed[1])
		switch dir {
		case "down":
			aim += moveCount
		case "up":
			aim -= moveCount
		case "forward":
			hPos += moveCount
			depth += moveCount * aim
		}
	}
	return hPos * depth
}
