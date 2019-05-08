package math

// this file deals with a term
// math var expression var
// simplify reevaluate, simplify, reevaluate
// expression is going to be a composition of terms
type Term interface {
	// if the var is alone then the nubmber is zero
	Setnumber() Term
	Setvars() Term
	SetExponent() Term
	GetTerm() Term
}

type Mxvar struct {
	Xvar bool
}

type Myvar struct {
	Yvar bool
}

// type Exponent struct {
// 	// if there is no exponent set this value to 1
// 	Ex int
// }

type Mvar struct {
	Mxvar
	Myvar
}

// term product
type TermP struct {
	Number   int
	Mvar     // check if one or the other is not nil
	Exponent int
}

// type TermBuilder struct {
// 	// math term
// 	mt TermProduct
// }

func (t *TermP) Setnumber(n int) *TermP {
	t.Number = n
	return t
}
func (t *TermP) SetExponent(n int) *TermP {
	t.Exponent = n
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
