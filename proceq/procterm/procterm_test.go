package procterm

import (
	"fmt"
	"testing"
)

func TestGetLopertrs(t *testing.T) {

	t1 := []string{`+`}
	t2 := []string{`+`, "-"}

	var tests = []struct {
		input string
		want  []string
	}{
		{`2x**2+3y**2=15`, t1},
		{`3x**2+2x+1=y`, t2},
	}
	for i, test := range tests {
		if got := procTerm(test.input); got[i] != test.want[i] {
			fmt.Printf("want is: %v, 	got is:%v", test.want, got)
			t.Errorf("Input Equation(%v) = %v", test.input, got)
		}
		i++
	}

}
