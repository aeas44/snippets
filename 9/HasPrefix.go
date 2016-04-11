package main

import "fmt"

// HasPrefix is just to solve 計算機科学一問一答
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix is just to solve 計算機科学一問一答
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains is just to solve 計算機科学一問一答
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func main() {
	s := "Hello, world"
	hello := "Hello"
	world := "world"

	fmt.Println("HasPrefix(), expecting: true, false |", HasPrefix(s, hello), HasPrefix(s, world))
	fmt.Println("HasSuffix(), expecting: false, true |", HasSuffix(s, hello), HasSuffix(s, world))
	fmt.Println("Contains(), expecting: true, true, false |", Contains(s, hello), Contains(s, world), Contains(s, "something"))
}
