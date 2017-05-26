package proceq

import (
	"fmt"
	"testing"
)

// func TestHsfrst(t *testing.T) {
// 	tv := make([]string, 0, 20)
// 	txt := make([]string, 0, 20)
// 	txt = append(txt, `2x**2+3y**2`, "15")

// 	tv = hsfrst(txt)
// 	if tv[0] != `2x**2+3y**2` {
// 		t.Error("Expected 2x**2+3y**2, got", tv[0])
// 	}
// 	if tv[1] != "15" {
// 		t.Error("expected 15, got", tv[1])
// 	}

// 	txt[0] = "15"
// 	txt[1] = `2x**2+3y**2`

// 	tv = hsfrst(txt)
// 	if tv[0] != `2x**2+3y**2` {
// 		t.Error("Expected 2x**2+3y**2, got", tv[0])
// 	}
// 	if tv[1] != "15" {
// 		t.Error("expected 15, got", tv[1])
// 	}

func TestHsfrst(t *testing.T) {

	t1 := []string{`2x**2+3y**2`, "15"}
	t2 := []string{"15", `2x**2+3y**2`}

	var tests = []struct {
		input []string
		want  []string
	}{
		{t1, t1},
		{t2, t1},
	}
	for i, test := range tests {
		if got := hsFrst(test.input); got[i] != test.want[i] {
			t.Errorf("Input Equation(%v) = %v", test.input, got)
		}
		i++
	}

}

func TestSepEx(t *testing.T) {

	t1 := []string{`2x**2+3y**2`, "15"}
	t2 := []string{`3x**2+2x+1`, "y"}

	var tests = []struct {
		input string
		want  []string
	}{
		{`2x**2+3y**2=15`, t1},
		{`3x**2+2x+1=y`, t2},
	}
	for i, test := range tests {
		if got := sepEx(test.input); got[i] != test.want[i] {
			fmt.Printf("want is: %v, 	got is:%v", test.want, got)
			t.Errorf("Input Equation(%v) = %v", test.input, got)
		}
		i++
	}
}

func TestSepTerms(t *testing.T) {

	t1 := []string{`2x**2`, "3y**2"}
	t2 := []string{`3x**2`, "2x", "1"}

	var tests = []struct {
		input string
		want  []string
	}{
		{`2x**2+3y**2`, t1},
		{`3x**2+2x+1`, t2},
	}
	for i, test := range tests {
		if got := sepTerms(test.input); got[i] != test.want[i] {
			fmt.Printf("want is: %v, 	got is:%v", test.want, got)
			t.Errorf("Input Equation(%v) = %v", test.input, got)
		}
		i++
	}

}

func TestGetLopertrs(t *testing.T) {

	t1 := []string{`+`}
	t2 := []string{`+`, "-"}

	var tests = []struct {
		input string
		want  []string
	}{
		{`2x**2+3y**2`, t1},
		{`3x**2+2x-1`, t2},
	}
	for i, test := range tests {
		if got := getLopertrs(test.input); got[i] != test.want[i] {
			fmt.Printf("want is: %v, 	got is:%v", test.want, got)
			t.Errorf("Input Equation(%v) = %v", test.input, got)
		}
		i++
	}

}

func procEq(t *testing.T) {

	t1 := []string{`+`}
	t2 := []string{`+`, "-"}

	var tests = []struct {
		input string
		want  []string
	}{
		{`2x**2+3y**2`, t1},
		{`3x**2+2x-1`, t2},
	}
	for i, test := range tests {
		if got := getLopertrs(test.input); got[i] != test.want[i] {
			fmt.Printf("want is: %v, 	got is:%v", test.want, got)
			t.Errorf("Input Equation(%v) = %v", test.input, got)
		}
		i++
	}

}
