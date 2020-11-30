package AoC2020

import "testing"

func TestHello(t *testing.T){
	want := "Hello Jim"
	if got := Hello(); got != want {
		t.Errorf("Hello() gave %q, wants %q", got, want)
	}
}