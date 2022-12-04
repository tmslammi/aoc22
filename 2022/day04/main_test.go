package day01

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	want := 602
	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}
}

func TestPar2(t *testing.T) {
	got := part2()
	want := 891
	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}
}
