package day01

import (
	"github.com/tmslammi/go-aoc/util"
	"time"
)

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 1 part 1")
	fileContent := util.ReadFile("input.txt")
	var toSum []int
	var biggestSum int
	for _, row := range fileContent {
		if row == "" {
			sum := 0
			for _, r := range toSum {
				sum += r
			}
			if sum > biggestSum {
				biggestSum = sum
			}
			toSum = make([]int, 0)
		} else {
			toSum = append(toSum, util.ToInt(row))
		}
	}
	return biggestSum
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 1 part 1")
	fileContent := util.ReadFile("input.txt")
	var toSum []int
	var biggestSums []int
	for _, row := range fileContent {
		if row == "" {
			sum := 0
			for _, r := range toSum {
				sum += r
			}
			biggestSums = append(biggestSums, sum)
			toSum = make([]int, 0)
		} else {
			toSum = append(toSum, util.ToInt(row))
		}
	}
	top := util.SortAsc(biggestSums)[:3]
	sum := util.Sum(top)
	return sum
}
