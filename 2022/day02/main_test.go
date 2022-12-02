package day01

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	if got != 8933 {
		t.Errorf("got = %d; want 8933", got)
	}
}

func TestPar2(t *testing.T) {
	got := part2()
	if got != 11998 {
		t.Errorf("got = %d; want 11998", got)
	}
}
