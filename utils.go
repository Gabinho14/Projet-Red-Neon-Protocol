package main

import (
	"fmt"
	"strings"
)

func readText() string {
	var text string
	fmt.Scan(&text)
	return text
}
func readInt() int {
	text := readText()
	if text == "" {
		return -1
	}
	number := 0
	for _, r := range text {
		if r < '0' || r > '9' {
			return -1
		}
		number = number*10 + int(r-'0')
	}
	return number
}
func clamp(value, low, high int) int {
	if value < low {
		return low
	}
	if value > high {
		return high
	}
	return value
}
func has(list []string, value string) bool {
	for _, x := range list {
		if x == value {
			return true
		}
	}
	return false
}
func letters(text string) bool {
	if text == "" {
		return false
	}
	for _, r := range text {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}
func formatName(name string) string {
	return strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
}
