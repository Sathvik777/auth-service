package util

import (
	"testing"
)

type stringTestPair struct {
	value          string
	expectedResult bool
}

var timestampTests = []stringTestPair{
	{"2017-01-01 15:00:00", true},
	{"2017-12-22 20:30:20", true},
	{"2017-12-22 20:34:00", true},
	{"2017-09-14 10:00:00", true},
	{"2017-12-22", false},
	{"2017-12-2215:30", false},
	{"-1, 1", false},
	{"notatime", false},
}

func TestValidityTimestamp(t *testing.T) {
	for _, pair := range timestampTests {
		res := IsStringTimestamp(pair.value)
		if res != pair.expectedResult {
			t.Error(
				"For", pair.value,
				"expected", pair.expectedResult,
				"got", res,
			)
		}
	}
}

var stringInSliceTests = []stringTestPair{
	{"ishere", true},
	{"isalsohere", true},
	{"thisishere", true},
	{"isnothere", false},
	{"-1, 1", false},
	{"notastringorisit", false},
}

func TestStringInSlice(t *testing.T) {
	stringArray := []string{"ishere", "thisishere", "isalsohere"}
	for _, pair := range stringInSliceTests {
		res := StringInSlice(pair.value, stringArray)
		if res != pair.expectedResult {
			t.Error(
				"For", pair.value,
				"expected", pair.expectedResult,
				"got", res,
			)
		}
	}
}
