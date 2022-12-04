package day01

import (
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
	}
	return result
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 4 part 2")
	input := util.ReadLines("input.txt")
	result := 0
	for _, row := range input {
		ranges := strings.Split(row, ",")
		r1 := buildRange(ranges[0])
		r2 := buildRange(ranges[1])
		if intersectsPartially(r1, r2) || intersectsPartially(r2, r1) {
			result++
		}
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
		if !util.ContainsInt(s2, i) {
			return false
		}
	}
	return true
}

func intersectsPartially(s1 []int, s2 []int) bool {
	for _, i := range s1 {
		if util.ContainsInt(s2, i) {
			return true
		}
	}
	return false
}
