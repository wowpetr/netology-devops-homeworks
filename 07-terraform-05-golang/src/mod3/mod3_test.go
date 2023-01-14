package main

import (
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestListModBy3Standard(t *testing.T) {
	start := 1
	end := 100
	want := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99}
	res, err := ListModBy3(start, end)
	if !Equal(res, want) || err != nil {
		t.Fatalf(`ListModBy3(%d, %d) = %v, %v, want match for %v, nil`, start, end, res, err, want)
	}
}

func TestListModBy3IncorrectBounds(t *testing.T) {
	start := 101
	end := 100
	res, err := ListModBy3(start, end)
	if res != nil || err == nil {
		t.Fatalf(`listModBy3(%d, %d) = %v, %v, wants [], error`, start, end, res, err)
	}
}
