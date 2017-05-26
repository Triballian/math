package main

import "fmt"

//import "unicode"

func main() {
	aslice := []string{"a", "batt", "is", " ", "on", "the", "kitchen", "floor"}
	fmt.Printf("%c", aslice[2][1])
	//pslice(aslice)

}

//func pslice(myslice []string) {
//	println(pslice)
//	for k, v := range myslice {
//		if unicode.IsSpace(v) {
//			continue
//		}
//		myslice = append(myslice[:k], myslice[k+1:]...)
//
//	}
//	println(pslice)
//}
