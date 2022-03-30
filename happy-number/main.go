package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var n int
	fmt.Print("Input a number: ")
	if _, err := fmt.Scanf("%d", &n); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %t", isHappy(n))
}

func isHappy(num int) bool {
	seen := make(map[int]struct{}, 0)
	for {
		if num == 1 {
			return true
		}
		if _, ok := seen[num]; ok {
			return false
		}
		seen[num] = struct{}{}
		num = sumDigitSquares(num)
	}
}

func sumDigitSquares(num int) int {
	numStr := strconv.Itoa(num)
	var acc int
	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))
		acc += digit * digit
	}
	return acc
}
