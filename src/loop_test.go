package main

import "testing"

func TestSum(t *testing.T) {
	str := loop()
	correct := "CodeEducation Rocks."
	if str != correct {
		t.Errorf("Loop return was incorrect, got: %v, want: %v.", str, correct)
	}
}
