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
	fmt.Printf("Output: %s", intToWord(n))
}

func intToWord(i int) string {
	if i == 0 {
		return "Zero"
	}
	var sb strings.Builder
	thousands := []string{ "", "Thousand", "Million", "Billion", "Trillion", "Quadrillion"}
	var idx int
	for i > 0 {
		if remainder := i % 1000 ; remainder != 0 {
			sb.WriteString(fmt.Sprintf("%s%s",helper(remainder), thousands[idx]))
		}
		idx++
		i /=1000
	}
	return strings.TrimRight(sb.String(), " ")
}


func helper(i int) string {
	if i == 0 {
		return ""
	}
	lessThanTwenty := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{ "", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	if i < 20 {
		return fmt.Sprintf("%s ",lessThanTwenty[i])
	} else if i < 100 {
		return fmt.Sprintf("%s %s ",tens[i/10], helper(i %10))
	} else {
		return fmt.Sprintf("%s Hundered %s ",lessThanTwenty[i/100],helper(i % 100))
	}
}