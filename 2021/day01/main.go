package day01

import (
	"github.com/tmslammi/go-aoc/util"
	"time"
)

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 1 part 1")
	fileContent := util.ReadLines("./input.txt")
	previousNumber := 0
	result := 0
	for i := 1; i < len(fileContent); i++ {
		current := util.ToInt(fileContent[i])
		if current > previousNumber {
			result++
		}
		previousNumber = current
	}
	return result
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 1 part 2")
	fileContent := util.ReadLines("./input.txt")
	previousSum := 0
	result := 0
	for i := 1; i < len(fileContent); i++ {
		end := i + 3
		if len(fileContent) < end {
			end = len(fileContent)
		}
		currentSlice := fileContent[i:end]
		sum := 0
		for _, row := range currentSlice {
			sum += util.ToInt(row)
		}
		if sum > previousSum {
			result++
		}
		previousSum = sum
	}
	return result
}
