package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Input string"
	result := toUpper(s)
	fmt.Println(s)
	fmt.Println(result)
}

func toUpper(in string) string {
	in = strings.ToUpper(in)
	return in
}
