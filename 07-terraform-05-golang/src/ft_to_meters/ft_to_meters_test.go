package main

import "testing"

func TestFeetToMeters(t *testing.T) {
	meters := 1.
	want := 0.3048
	val := FeetToMeters(meters)
	if val != want {
		t.Fatalf(`FeetToMeters(%v) = %v, want %v`, meters, val, want)
	}
}
