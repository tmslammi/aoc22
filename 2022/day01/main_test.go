package day01

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	if got != 0 {
		t.Errorf("part1() = %d; want 0", got)
	}
}
