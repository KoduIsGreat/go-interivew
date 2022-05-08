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
			RomanNumeral: "MCDLVII",
			Integer: 1457,
		},
		{
			RomanNumeral: "MCMXCIV",
			Integer:      1994,
		},
		{
			RomanNumeral: "MMDXXXIV",
			Integer: 2534,
		},
	} {
		t.Run(tc.RomanNumeral, func(t *testing.T) {
			if got := intToRoman(tc.Integer); got != tc.RomanNumeral {
				t.Fatalf("got %s, want %s", got, tc.RomanNumeral)
			}
		})
	}
}
