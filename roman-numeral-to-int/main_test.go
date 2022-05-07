package main

import "testing"

func TestRomanToInt(t *testing.T) {
	for _, tc := range []struct {
		RomanNumeral string
		Integer      int
	}{
		{
			RomanNumeral: "III",
			Integer:      3,
		},
		{
			RomanNumeral: "IV",
			Integer:      4,
		},
		{
			RomanNumeral: "V",
			Integer:      5,
		},
		{
			RomanNumeral: "VI",
			Integer:      6,
		},
		{
			RomanNumeral: "IX",
			Integer:      9,
		},
		{
			RomanNumeral: "X",
			Integer:      10,
		},
		{
			RomanNumeral: "L",
			Integer:      50,
		},
		{
			RomanNumeral: "XL",
			Integer:      40,
		},
		{
			RomanNumeral: "C",
			Integer:      100,
		},
		{
			RomanNumeral: "XC",
			Integer:      90,
		},
		{
			RomanNumeral: "MCMXCIV",
			Integer:      1994,
		},
	} {
		t.Run(tc.RomanNumeral, func(t *testing.T) {
			if got := romanToInt(tc.RomanNumeral); got != tc.Integer {
				t.Fatalf("got %d, want %d", got, tc.Integer)
			}
		})
	}
}
