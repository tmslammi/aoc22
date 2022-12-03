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

func findMatches(s1 string, s2 string) []string {
	set := hashset.New()
	firstSplit := strings.Split(s1, "")
	for _, r := range strings.Split(s2, "") {
		if util.Contains(firstSplit, r) {
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

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 3 part 2")
	return 0
}
