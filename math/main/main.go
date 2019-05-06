package main

import (
	"fmt"

	"github.com/Triballian/math/math"
)

func main() {
	thisvar := &math.TermP{}
	thisvar.Setnumber(4)
	thisvar.Setvars("x", true)
	fmt.Printf("thisvar.yvar:%t", thisvar.Yvar)
}
