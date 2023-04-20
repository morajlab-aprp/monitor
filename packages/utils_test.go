package main

import "testing"

func TestCreateRegex(t *testing.T) {
	regex := CreateRegex(`Test:\s*(\w|\d)+`, "i")

	if regex.Value != `Test:\s*(\w|\d)+` {
		t.Errorf("Regex value doesn't match %s", `Test:\s*(\w|\d)+`)
	}
}
