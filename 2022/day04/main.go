package day01

import (
	"github.com/samber/lo"
	"github.com/tmslammi/go-aoc/util"
	"strings"
	"time"
)

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 4 part 1")
	input := util.ReadLines("input.txt")
	result := 0
	for _, row := range input {
		ranges := strings.Split(row, ",")
		r1 := buildRange(ranges[0])
		r2 := buildRange(ranges[1])
		if intersectsFull(r1, r2) || intersectsFull(r2, r1) {
			result++
		}
		print(r1)
		print(r2)
	}
	return result
}

func buildRange(s string) []int {
	split := strings.Split(s, "-")
	start := util.ToInt(split[0])
	end := util.ToInt(split[1])
	r := util.BuildRange(start, end)
	return r
}

func intersectsFull(s1 []int, s2 []int) bool {
	for _, i := range s1 {
		if !lo.Contains(s2, i) {
			return false
		}
	}
	return true
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 4 part 2")
	return 0
}
