package main

//need to identify absolute value
// need to count positions and need to note the highest position
// need to note if the expression has an absolute value, if so , in what possitions -
// is the abs including
// number of terms should be seperate from number of postions
import (
	"fmt"

	"github.com/Triballian/math/math"
)

func main() {
	// this file is for test purposes only
	thisvar := &math.TermP{}
	thisvar.Setnumber(4)
	thisvar.Setvars("x", true)
	fmt.Printf("thisvar.yvar:%t\n", thisvar.Yvar)
	// Expressions
	manufacturingComplex := math.ManufacturingDirector{}
	monomialBuilder := &math.MonomialBuilder{}
	manufacturingComplex.SetBuilder(monomialBuilder)
	manufacturingComplex.Construct(1, 0, 1)
	monomial := monomialBuilder.GetMathExpression()
	fmt.Println("this is monomial:%V", monomial)

}
