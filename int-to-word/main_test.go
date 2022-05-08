package main

import "testing"

func TestRomanToInt(t *testing.T) {
	for _, tc := range []struct {
		Word    string
		Integer int
	}{
		{
			Word:    "Three",
			Integer: 3,
		},
		{
			Word: "Zero",
			Integer: 0,
		},
		{
			Word: "One Hundered Fifty One",
			Integer: 151,
		},
		{
			Integer: 12345,
			Word: "Twelve Thousand Three Hundred Forty Five",
		},

	} {
		t.Run(tc.Word, func(t *testing.T) {
			if got := intToWord(tc.Integer); got != tc.Word {
				t.Fatalf("got %s, want %s", got, tc.Word)
			}
		})
	}
}
