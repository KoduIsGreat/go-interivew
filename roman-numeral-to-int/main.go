package main

import (
	"fmt"
	"log"
)

func main() {
	var n string
	fmt.Print("Input a Roman Numeral: ")
	if _, err := fmt.Scanf("%s", &n); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %d", romanToInt(n))
}

var lookup = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	var sum, max int
	for cur := len(s) - 1; cur != -1; cur-- {
		if num := lookup[rune(s[cur])]; num >= max {
			sum += num
			max = num
		} else {
			sum -= num
		}
	}
	return sum
}
