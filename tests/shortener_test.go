package tests

import (
	"fmt"
	"testing"

	"github.com/ckng0221/url-shortener/utils"
)

func Test_ConvertIntegerToBase62(t *testing.T) {
	testCases := []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{36, "A"},
		{61, "Z"},
		{62, "10"},
		{63, "11"},
		{140716, "ABC"},
	}

	for _, tc := range testCases {

		// run can see individual in verbose
		t.Run(fmt.Sprintf("in: %v", tc.in), func(t *testing.T) {

			got := utils.ConvertIntegerToBase62(tc.in)
			if got != tc.out {
				t.Errorf("Failed, except %s, got %s", tc.out, got)
			}
		})
	}
}

func Test_ConvertIntegerToBase62WithSwitch(t *testing.T) {
	testCases := []struct {
		in  int
		out string
	}{
		{1, "1"},
		{2, "2"},
		{36, "A"},
		{61, "Z"},
		{62, "10"},
		{63, "11"},
		{140716, "ABC"},
	}

	for _, tc := range testCases {

		// run can see individual in verbose
		t.Run(fmt.Sprintf("in: %v", tc.in), func(t *testing.T) {

			got := utils.ConvertIntegerToBase62WithMap(tc.in)
			if got != tc.out {
				t.Errorf("Failed, except %s, got %s", tc.out, got)
			}
		})
	}
}

func Test_ConvertBase62ToInteger(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{"1", 1},
		{"A", 36},
		{"10", 62},
		{"11", 63},
		{"ABC", 140716},
	}

	for _, tc := range testCases {

		// run can see individual in verbose
		t.Run(fmt.Sprintf("in: %v", tc.in), func(t *testing.T) {
			got := utils.ConvertBase62ToInteger(tc.in)
			if got != tc.out {
				t.Errorf("Failed, except %v, got %v", tc.out, got)
			}
		})
	}
}
