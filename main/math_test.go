package main

import (
	"fmt"
	"testing"

	"github.com/Triballian/math/values"
)

func TestSets(t *testing.T) {
	// d := 0

	fsets := []values.Fset{}
	fset := values.Fset{}
	fset.Domain = 1
	fsets = append(fsets, fset)
	fsets[0].Domain = 1

	if fsets[0].Domain != 0 {
		t.Errorf(`ftest.Domain is: %v`, fsets[0].Domain)
	}

}
func TestValues(t *testing.T) {
	fmt.Println(values.Equestion{Isreq: true, Q: "This is Eqeustion"})

}
