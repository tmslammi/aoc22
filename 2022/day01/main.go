package day01

import (
	"github.com/tmslammi/go-aoc/util"
	"time"
)

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 1 part 1")
	fileContent := util.ReadLines("input.txt")
	var biggestSum int
	grouped := util.ChunkByChar(fileContent, "")
	for _, row := range grouped {
		sum := util.Sum(util.SliceToInt(row))
		if sum > biggestSum {
			biggestSum = sum
		}
	}
	return biggestSum
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 1 part 2")
	fileContent := util.ReadLines("input.txt")
	var biggestSums []int
	grouped := util.ChunkByChar(fileContent, "")
	for _, row := range grouped {
		sum := util.Sum(util.SliceToInt(row))
		biggestSums = append(biggestSums, sum)
	}
	top := util.SortAsc(biggestSums)[:3]
	return util.Sum(top)
}
