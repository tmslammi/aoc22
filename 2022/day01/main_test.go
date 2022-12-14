package day01

import "testing"

func TestPart1(t *testing.T) {
	got := part1()
	want := 71124
	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}
}

func TestPar2(t *testing.T) {
	got := part2()
	want := 204639
	if got != want {
		t.Errorf("got = %d; want %d", got, want)
	}
}
