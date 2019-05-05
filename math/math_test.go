package math

import (
	"testing"
)

func testmvar(t *testing.T) {
	expectednumber := 4
	expectedmxvar := true
	expectedmyvar := true
	thisvar := term{}
	thisvar.number = 4
	thisvar.xvar = true
	thisvar.yvar = true
	if thisvar.number != expectednumber {
		t.Errorf("The expecteed number:%d does not match thisvar.number:%d", expectednumber, thisvar.number)
	}
	if thisvar.xvar != expectedmxvar {
		t.Errorf("The expected xvar bool:%t does not match thisvar.xvar:%t", expectedmxvar, thisvar.xvar)
	}
	if thisvar.yvar != expectedmyvar {
		t.Errorf("The Expedted  yvarbool:%t does not match thisvar.yvar:%t", expectedmyvar, thisvar.yvar)
	}

}
