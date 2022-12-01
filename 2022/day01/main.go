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
			sum := util.Sum(toSum)
			if sum > biggestSum {
				biggestSum = sum
			}
			toSum = nil
		} else {
			toSum = append(toSum, util.ToInt(row))
		}
	}
	return biggestSum
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 1 part 2")
	fileContent := util.ReadFile("input.txt")
	var toSum []int
	var biggestSums []int
	for _, row := range fileContent {
		if row == "" {
			sum := util.Sum(toSum)
			biggestSums = append(biggestSums, sum)
			toSum = nil
		} else {
			toSum = append(toSum, util.ToInt(row))
		}
	}
	top := util.SortAsc(biggestSums)[:3]
	return util.Sum(top)
}
