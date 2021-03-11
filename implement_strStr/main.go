package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack, needle string) int {
	h := len(haystack)
	n := len(needle)
	if n == 0 {
		return 0
	}
	i := 0
	for i <= h - n {
		start := 0
		for start < n && haystack[i+start] == needle[start] {
			start++
		}
		if start == n {
			return i
		}
		i++
	}
	return -1
}
