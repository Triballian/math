package hsfrst

import "testing"

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
	t1 := make([]string, 1, 20)
	t2 := make([]string, 1, 20)

	t1 = append(t1, `2x**2+3y**2`, "15")
	t2 = append(t2, "15", `2x**2+3y**2`)

	var tests = []struct {
		input []string
		want  []string
	}{
		{t1, t1},
		{t2, t1},
	}
	for i, test := range tests {
		if got := hsfrst(test.input); got[i] != test.want[i] {
			t.Errorf("Input Equation(%v) = %v", got, test.want)
		}
		i++
	}

}
