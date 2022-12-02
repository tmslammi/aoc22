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
		enemyScore := resolveDrawScore(draws[0])
		playerScore := resolveDrawScore(draws[1])

		score := getScore(enemyScore, playerScore)
		total += score + playerScore
	}

	return total
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 2 part 2")
	input := util.ReadLines("input.txt")

	total := 0

	for _, row := range input {
		draws := strings.Split(row, " ")
		enemyScore := resolveDrawScore(draws[0])
		playerScore := resolvePlayerMove(enemyScore, resolveWhatPlayerShouldDo(draws[1]))

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

func resolveDrawScore(s string) int {
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

func resolveWhatPlayerShouldDo(s string) int {
	switch s {
	case "X":
		return LOSS
	case "Y":
		return DRAW
	case "Z":
		return WIN
	}

	return 0
}

func resolvePlayerMove(enemy int, player int) int {
	if enemy == ROCK && player == DRAW {
		return ROCK
	}
	if enemy == ROCK && player == LOSS {
		return SCISSORS
	}
	if enemy == ROCK && player == WIN {
		return PAPER
	}
	if enemy == PAPER && player == DRAW {
		return PAPER
	}
	if enemy == PAPER && player == LOSS {
		return ROCK
	}
	if enemy == PAPER && player == WIN {
		return SCISSORS
	}
	if enemy == SCISSORS && player == DRAW {
		return SCISSORS
	}
	if enemy == SCISSORS && player == LOSS {
		return PAPER
	}
	if enemy == SCISSORS && player == WIN {
		return ROCK
	}
	return 0
}
