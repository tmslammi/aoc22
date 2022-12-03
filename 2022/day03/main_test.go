package day01

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	if got != 7917 {
		t.Errorf("got = %d; want 7917", got)
	}
}

func TestPar2(t *testing.T) {
	got := part2()
	if got != 2585 {
		t.Errorf("got = %d; want 2585", got)
	}
}
