package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var n int
	fmt.Print("Input a Roman Numeral: ")
	if _, err := fmt.Scanf("%d", &n); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %s", intToRoman(n))
}

func intToRoman(i int) string {
	var sb strings.Builder
	var nums = []int {1000,900,500,400,100,90,50,40,10,9,5,4,1}
	var symbols = []string { "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV","I" }
	for idx, num := range nums {
		for i >= num {
			sb.WriteString(symbols[idx])
			i -= num
		}
	}
	return sb.String()
}
