package day02

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	if got != 1882980 {
		t.Errorf("part1() = %d; want 1882980", got)
	}
}

func TestPart2(t *testing.T) {
	got := part2()
	if got != 1971232560 {
		t.Errorf("part2() = %d; want 1971232560", got)
	}
}
