package day01

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/tmslammi/go-aoc/util"
	"strings"
	"time"
)

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 3 part 1")
	input := util.ReadLines("input.txt")
	total := 0
	for _, row := range input {
		first := row[0 : len(row)/2]
		second := row[len(row)/2:]
		matches := findMatches(first, second)
		total += calculateScoreForMatches(matches)
	}
	return total
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 3 part 2")
	input := util.ReadLines("input.txt")
	total := 0
	for i := 0; i < len(input); i += 3 {
		chunk := input[i : i+3]
		matches := findMatchesBatch(chunk)
		total += calculateScoreForMatches(matches)
	}
	return total
}

func findMatches(s1 string, s2 string) []string {
	set := hashset.New()
	firstSplit := strings.Split(s1, "")
	for _, r := range strings.Split(s2, "") {
		if util.ContainsString(firstSplit, r) {
			set.Add(fmt.Sprint(r))
		}
	}
	return util.InterfaceToSlice(set.Values())
}

func findMatchesBatch(slice []string) []string {
	set := hashset.New()
	firstSplit := strings.Split(slice[0], "")
	secondSplit := strings.Split(slice[1], "")
	for _, r := range strings.Split(slice[2], "") {
		if util.ContainsString(firstSplit, r) && util.ContainsString(secondSplit, r) {
			set.Add(fmt.Sprint(r))
		}
	}
	return util.InterfaceToSlice(set.Values())
}

func calculateScoreForMatches(matches []string) int {
	score := 0
	for _, m := range matches {
		ascii := int(m[0])
		if strings.ToUpper(m) == m {
			score += ascii - 38
		} else {
			score += ascii - 96
		}
	}
	return score
}
