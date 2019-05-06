package math

import (
	"testing"
)

func testmvar(t *testing.T) {
	expectednumber := 4
	expectedmxvar := true
	expectedmyvar := true
	//manufacturingComplex := ManufacturingDirector{}

	thisvar := &TermP{}
	thisvar.Setnumber(4)
	thisvar.Setvars("x", true).Setvars("y", true)
	if thisvar.Number != expectednumber {
		t.Errorf("The expecteed number:%d does not match thisvar.number:%d", expectednumber, thisvar.Number)
	}
	if thisvar.Xvar != expectedmxvar {
		t.Errorf("The expected xvar bool:%t does not match thisvar.xvar:%t", expectedmxvar, thisvar.Xvar)
	}
	if thisvar.Yvar != expectedmyvar {
		t.Errorf("The Expedted  yvarbool:%t does not match thisvar.yvar:%t", expectedmyvar, thisvar.Yvar)
	}

}
