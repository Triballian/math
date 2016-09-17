package main

import "unicode"

func main() {
	aslice := []string{"a", "batt", "is", " ", "on", "the", "kitchen", "floor"}
	pslice(aslice)

}
func pslice(myslice []string) {
	println(pslice)
	for k, v := range myslice {
		if unicode.IsSpace(v) {
			continue
		}
		myslice = append(myslice[:k], myslice[k+1:]...)

	}
	println(pslice)
}
