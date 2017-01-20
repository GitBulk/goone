package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("-- Collection --")
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(input string) bool {
		return strings.HasPrefix(input, "p")
	}))

	fmt.Println(All(strs, func(input string) bool {
		return strings.HasPrefix(input, "p")
	}))

	fmt.Println(Filter(strs, func(input string) bool {
		return strings.Contains(input, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}

func Index(vs []string, search string) int {
	for i, item := range vs {
		if item == search {
			return i
		}
	}
	return -1
}

func Include(vs []string, search string) bool {
	return Index(vs, search) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, item := range vs {
		if f(item) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, item := range vs {
		if !f(item) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	filterd := make([]string, 0)
	for _, item := range vs {
		if f(item) {
			filterd = append(filterd, item)
		}
	}
	return filterd
}

func Map(vs []string, f func(string) string) []string {
	dest := make([]string, len(vs))
	for i, item := range vs {
		dest[i] = f(item)
	}
	return dest
}
