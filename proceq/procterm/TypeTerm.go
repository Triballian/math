package procterm

import "strings"

func TypeTerm(t string, v string) ([]string, bool) {
	r := strings.Split(t, v)
	if len(r) < 2 {
		return r, true
	}
	return r, false

}
