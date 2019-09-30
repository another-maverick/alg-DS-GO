package main

import "testing"

type eachEntry struct{
	actual string
	length int
	expected string
}

func newEachEntry(actual string, length int, expected string) *eachEntry {
	return &eachEntry{actual: actual, length: length, expected: expected}
}

func TestURLify(t *testing.T) {
	data := newEachEntry("much ado about nothing      ", 22, "much%20ado%20about%20nothing" )
	result := URLify(data.actual, data.length)
	if result != data.expected{
		t.Errorf("We got %s and we expected %s", result, data.expected)
	}
}
