package proceq

import (
	"fmt"
	"strings"
	"unicode"
)

// sepEx seperates expressoins in an equation
func sepEx(t string) []string {
	eq := make([]string, 0, 20)
	var fstring string
	// remove space and convert text into a slice
	for _, value := range t {
		if unicode.IsSpace(value) {
			continue
		}
		fstring = fmt.Sprintf("%c", value)
		// fmt.Printf("%c", value)
		eq = append(eq, fstring)

	}

	fstring = strings.Join(eq, "")

	eqEx := strings.Split(fstring, `=`)
	// fmt.Printf("eq is:%v", fstring)
	return eqEx
}
