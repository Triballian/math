package math

// math var expression var
// simplify reevaluate, simplify, reevaluate
type Term interface {
	// if the var is alone then the nubmber is zero
	Setnumber() Term
	Setvars() Term
	GetTerm() Term
}

// Director
type ManufacturingDirector struct {
	Builder Term
}

// construct method
func (f *ManufacturingDirector) Construct() {
	f.Builder.Setnumber().Setvars()
}

type Mxvar struct {
	Xvar bool
}

type Myvar struct {
	Yvar bool
}

type Mvar struct {
	Mxvar
	Myvar
}

// term product
type TermP struct {
	Number int
	Mvar   // check if one or the other is not nil
}

// type TermBuilder struct {
// 	// math term
// 	mt TermProduct
// }

func (t *TermP) Setnumber(n int) *TermP {
	t.Number = n
	return t
}
func (t *TermP) Setvars(thisvar string, varval bool) *TermP {
	if thisvar == "x" || thisvar == "X" {
		t.Xvar = varval
	}
	if thisvar == "y" || thisvar == "Y" {
		t.Yvar = varval
	}

	return t
}

func (t *TermP) GetTerm(x bool) *TermP {
	return t
}
