package main

import "testing"

func TestMin(t *testing.T) {
	arr := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	want := 9
	min, err := Min(arr)
	if min != want || err != nil {
		t.Fatalf(`Min(%v) = %v, nil, want %v, %v`, arr, min, want, err)
	}
}

func TestMinEmpty(t *testing.T) {
	arr := []int{}
	want := 0
	min, err := Min(arr)
	if min != want || err == nil {
		t.Fatalf(`Min(%v) = %v, error, want %v, %v`, arr, min, want, err)
	}
}
