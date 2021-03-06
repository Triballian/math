package math

import (
	"fmt"
	"testing"
)

func TestMvar(t *testing.T) {

	expectednumber := 4
	expectedmxvar := true
	expectedmyvar := true
	// tmanufacturingComplex := TManufacturingDirector{}
	tmanufacturingComplex := GetTManInstance()
	constantBuilder := &ConstantTerm{}

	// thisvar := &TermP{}
	constantBuilder.TConstruct(0, 0, false, 1, 15, "")
	// constantBuilder.SetExponent(0)
	// constantBuilder.Setvars("x", true).Setvars("y", true)

	if constantBuilder.Number != expectednumber {
		t.Errorf("The expecteed number:%d does not match thisvar.number:%d", expectednumber, constantBuilder.Number)
	}
	if constantBuilder.Xvar != expectedmxvar {
		t.Errorf("The expected xvar bool:%t does not match thisvar.xvar:%t", expectedmxvar, constantBuilder.Xvar)
	}

	if thisvar.Yvar != expectedmyvar {
		t.Errorf("The Expected  yvarbool:%t does not match thisvar.yvar:%t", expectedmyvar, constantBuilder.Yvar)
	}

	manufacturingComplex := GetManInstance()
	monomialBuilder := &MonomialBuilder{}
	manufacturingComplex.SetBuilder(monomialBuilder)
	manufacturingComplex.Construct(1, 0, 1)
	monomial := monomialBuilder.GetMathExpression()
	expectedcount := 1
	expectedabssets := 0
	exppectedoperators := 1
	if monomial.TermCount != expectedcount {
		t.Errorf("The expected count:%d does not mathch monomial.TermCount:%d", expectedcount, monomial.TermCount)
	}
	if monomial.ABSsets != expectedabssets {
		t.Errorf("The expected count:%d does not mathch monomial.TermCount:%d", expectedabssets, monomial.ABSsets)
	}
	if monomial.Operators != exppectedoperators {
		t.Errorf("The expected count:%d does not mathch monomial.TermCount:%d", exppectedoperators, monomial.Operators)
	}

}

type testCase struct {
	mcount             int
	expectedmcount     int
	mabssets           int
	expectedmabssets   int
	moperators         int
	expectedmoperators int
}

func TestManyExpressions(t *testing.T) {
	manufacturingComplex := GetManInstance()
	monomialBuilder := &MonomialBuilder{}
	manufacturingComplex.SetBuilder(monomialBuilder)
	testCases := []testCase{
		{1, 1,
			1, 1,
			1, 1},
		{1, 1,
			0, 0,
			0, 0},
		{1, 1,
			1, 1,
			0, 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.mcount), func(t *testing.T) {
			manufacturingComplex.Construct(tc.mcount, tc.mabssets, tc.moperators)
			if tc.mcount != tc.expectedmcount {
				t.Fatalf("mcount: %d != %d", tc.mcount, tc.expectedmcount)
			}
			if tc.mabssets != tc.expectedmabssets {
				t.Fatalf("abssets: %d != %d", tc.mabssets, tc.expectedmabssets)
			}
			if tc.moperators != tc.expectedmoperators {
				t.Fatalf("moperators: %d != %d", tc.moperators, tc.expectedmoperators)
			}
		})
	}
}
