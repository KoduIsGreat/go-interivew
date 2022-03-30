package main

import "testing"

func FuzzIsHappy(f *testing.F) {
	f.Add(19, true)
	f.Add(2, false)
	f.Fuzz(func(t *testing.T, num int, want bool) {
		if got := isHappy(num); got != want {
			t.Fatalf("got %t , want %t for num %d", got, want, num)
		}
	})
}
