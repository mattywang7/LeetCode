package main

import "fmt"

// Given a string containing digits from 2-9 inclusive,
// return all possible letter combinations that the number could represent. Return the answer in any order.

// Input:	digits = "23"
// Output:	["ad","ae","af","bd","be","bf","cd","ce","cf"]

// Input:	digits = ""
// Output:	[]

// Input:	digits = "2"
// Output:	["a", "b", "c"]

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

// 0ms
// 2MB
func letterCombinations(digits string) []string {
	// In this case, have to use []string as value
	// if string as value, cannot use curStr + ch to get a new string quickly
	m := map[byte][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var res []string
	if digits == "" {
		return res
	}
	backtrack("", &res, 0, digits, m)
	return res
}

func backtrack(curStr string, res *[]string, index int, digits string, digitsMap map[byte][]string) {
	if index == len(digits) {
		*res = append(*res, curStr)
		return
	}
	for _, ch := range digitsMap[digits[index]] {
		backtrack(curStr+ch, res, index+1, digits, digitsMap)
	}
}
