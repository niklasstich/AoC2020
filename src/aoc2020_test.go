package AoC2020

import "testing"

func TestDay1_1(t *testing.T) {
	want := 41979
	if got := Day1_1(); got != want {
		t.Errorf("Day1_1() gave %q, wants %q", got, want)
	}
}