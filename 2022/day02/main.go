package day01

import (
	"github.com/tmslammi/go-aoc/util"
	"strings"
	"time"
)

const (
	A = ROCK
	B = PAPER
	C = SCISSORS
	X = ROCK
	Y = PAPER
	Z = SCISSORS
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

const (
	WIN  = 6
	DRAW = 3
	LOSS = 0
)

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 2 part 1")
	input := util.ReadLines("input.txt")

	total := 0

	for _, row := range input {
		draws := strings.Split(row, " ")
		enemyScore := resolveDrawScores(draws[0])
		playerScore := resolveDrawScores(draws[1])

		score := getScore(enemyScore, playerScore)
		total += score + playerScore
	}

	return total
}

func getScore(enemy int, player int) int {
	if enemy == player {
		return DRAW
	}
	if enemy == ROCK && player != PAPER {
		return LOSS
	}
	if enemy == PAPER && player != SCISSORS {
		return LOSS
	}
	if enemy == SCISSORS && player != ROCK {
		return LOSS
	}
	return WIN
}

func resolveDrawScores(s string) int {
	switch s {
	case "A":
		return A
	case "X":
		return X
	case "B":
		return B
	case "Y":
		return Y
	case "C":
		return C
	case "Z":
		return Z
	}

	return 0
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 2 part 2")
	return 0
}
