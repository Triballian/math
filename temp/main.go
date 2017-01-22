package main

import "fmt"

func hsfrst(s []string) []string {

	if len(s[0]) < len(s[1]) {
		s[1], s[0] = s[0], s[1]
	}
	return s
}

func main() {
	result := make([]string, 0, 20)
	text := make([]string, 0, 20)
	text = append(text, `2x**2+3y**2`, "15")
	// text[1] = "15"
	result = hsfrst(text)
	fmt.Println("result is %s", result)
}
