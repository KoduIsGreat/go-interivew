package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	fmt.Print("Input a number: ")
	if _, err := fmt.Scanf("%d", &n); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output: %t", isHappy(n))
}

func isHappyMap(num int) bool {
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

func isHappy(num int) bool {
	slow := sumDigitSquares(num)
	fast := sumDigitSquares(slow)
	for slow != fast && fast != 1 {
		slow = sumDigitSquares(slow)
		fast = sumDigitSquares(sumDigitSquares(fast))
	}
	return fast == 1
}

func sumDigitSquares(num int) int {
	var acc int
	for num != 0 {
		digit := num % 10
		acc += digit * digit
		num /= 10
	}
	return acc
}
