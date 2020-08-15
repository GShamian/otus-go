package main

import (
	"strings"
	"unicode"
)

func unpackString(s string) string {
	var result strings.Builder
	var currentRune rune
	for _, r := range s {
		if unicode.IsDigit(r) != true {
			result.WriteRune(r)
			currentRune = r
		} else if currentRune != 0 {
			for i := 1; i < int(r-'0'); i++ {
				result.WriteRune(currentRune)
			}
		}
	}
	return result.String()
}

func main() {
	//fmt.Println(unpackString("43"))
}
