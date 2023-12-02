package one_test

import (
	"testing"

	"github.com/piotrostr/aoc23/one"
)

func TestOne(t *testing.T) {
	input := `1abc2
	pqr3stu8vwx
	a1b2c3d4e5f
	treb7uchet`
	got := one.Trebuchet(input)
	want := 142
	if got != want {
		t.Errorf("Trebuchet(input) = %d; want %d", got, want)
	}
}
