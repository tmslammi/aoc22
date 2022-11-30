package day01

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	if got != 1400 {
		t.Errorf("part1() = %d; want 1400", got)
	}
}

func TestPart2(t *testing.T) {
	got := part2()
	if got != 1429 {
		t.Errorf("part1() = %d; want 1429", got)
	}
}
