package util

import "testing"

func TestChunkByChar(t *testing.T) {
	s := []string{"a", "a", "", "b", "b"}
	got := ChunkByChar(s, "")
	if len(got) != 2 {
		t.Errorf("got = %s; want 2", got)
	}
}

func TestFindMax(t *testing.T) {
	s := []int{50, 100, 20}
	got := FindMax(s)
	if got != 100 {
		t.Errorf("got = %d; want 100", got)
	}
}
