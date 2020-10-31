package leetcode

import "strings"

var table = []struct {
	symbol string
	value  int
}{
	{symbol: "M", value: 1000},
	{symbol: "CM", value: 900},
	{symbol: "D", value: 500},
	{symbol: "CD", value: 400},
	{symbol: "C", value: 100},
	{symbol: "XC", value: 90},
	{symbol: "L", value: 50},
	{symbol: "XL", value: 40},
	{symbol: "X", value: 10},
	{symbol: "IX", value: 9},
	{symbol: "V", value: 5},
	{symbol: "IV", value: 4},
	{symbol: "I", value: 1},
}

func intToRoman(num int) string {
	var roman strings.Builder

	n := len(table)
	for i := 0; i < n; i++ {
		for num >= table[i].value {
			num -= table[i].value
			roman.WriteString(table[i].symbol)
		}
	}

	return roman.String()
}
