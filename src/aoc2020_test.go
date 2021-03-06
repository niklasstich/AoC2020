package AoC2020

import "testing"

func TestDay1_1(t *testing.T) {
	want := 41979
	if got := Day1_1(); got != want {
		t.Errorf("Day1_1() gave %d, wants %d", got, want)
	}
}

func TestDay1_2(t *testing.T) {
	want := 193416912
	if got := Day1_2(); got != want {
		t.Errorf("Day1_2() gave %d, wants %d", got, want)
	}
}

func TestDay2_1(t *testing.T) {
	want := 434
	if got := Day2_1(); got != want {
		t.Errorf("Day2_1() gave %d, wants %d", got, want)
	}
}

func TestDay2_2(t *testing.T) {
	want := 509
	if got := Day2_2(); got != want {
		t.Errorf("Day2_2() gave %d, wants %d", got, want)
	}
}

func TestDay3_1(t *testing.T) {
	want := 178
	if got := Day3_1(); got != want {
		t.Errorf("Day3_1() gave %d, wants %d", got, want)
	}
}

func TestDay3_2(t *testing.T) {
	want := 3492520200
	if got := Day3_2(); got != want {
		t.Errorf("Day3_2() gave %d, wants %d", got, want)
	}
}

func TestDay4_1(t *testing.T) {
	want := 256
	if got := Day4_1(); got != want {
		t.Errorf("Day4_1() gave %d, wants %d", got, want)
	}
}

func TestDay4_2(t *testing.T) {
	want := 198
	if got := Day4_2(); got != want {
		t.Errorf("Day4_2() gave %d, wants %d", got, want)
	}
}

func TestDay5_1(t *testing.T) {
	want := 994
	if got := Day5_1(); got != want {
		t.Errorf("Day5_1() gave %d, wants %d", got, want)
	}
}

func TestDay5_2(t *testing.T) {
	want := 741
	if got := Day5_2(); got != want {
		t.Errorf("Day5_2() gave %d, wants %d", got, want)
	}
}
