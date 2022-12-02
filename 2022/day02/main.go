package day01

import (
	"github.com/tmslammi/go-aoc/util"
	"strings"
	"time"
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

var resultMap = map[[2]int]int{
	{ROCK, DRAW}:     ROCK,
	{ROCK, LOSS}:     SCISSORS,
	{ROCK, WIN}:      PAPER,
	{PAPER, DRAW}:    PAPER,
	{PAPER, LOSS}:    ROCK,
	{PAPER, WIN}:     SCISSORS,
	{SCISSORS, DRAW}: SCISSORS,
	{SCISSORS, LOSS}: PAPER,
	{SCISSORS, WIN}:  ROCK,
}

func part1() int {
	defer util.TimeTrack(time.Now(), "Day 2 part 1")
	input := util.ReadLines("input.txt")
	total := 0
	for _, row := range input {
		draws := strings.Split(row, " ")
		enemyWeaponValue := getWeaponValue(draws[0])
		playerWeaponValue := getWeaponValue(draws[1])
		score := getScore(enemyWeaponValue, playerWeaponValue)
		total += score + playerWeaponValue
	}
	return total
}

func part2() int {
	defer util.TimeTrack(time.Now(), "Day 2 part 2")
	input := util.ReadLines("input.txt")
	total := 0
	for _, row := range input {
		draws := strings.Split(row, " ")
		enemyWeaponValue := getWeaponValue(draws[0])
		playerWeaponValue := resolvePlayerWeapon(enemyWeaponValue, getExpectedRoundResult(draws[1]))
		score := getScore(enemyWeaponValue, playerWeaponValue)
		total += score + playerWeaponValue
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

func getWeaponValue(s string) int {
	switch s {
	case "A", "X":
		return ROCK
	case "B", "Y":
		return PAPER
	case "C", "Z":
		return SCISSORS
	}
	return 0
}

func getExpectedRoundResult(s string) int {
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

func resolvePlayerWeapon(enemy int, player int) int {
	pair := [2]int{enemy, player}
	return resultMap[pair]
}
