package day01

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	if got != 71124 {
		t.Errorf("got = %d; want 71124", got)
	}
}

func TestPar2(t *testing.T) {
	got := part2()
	if got != 204639 {
		t.Errorf("got = %d; want 204639", got)
	}
}
