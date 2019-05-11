// Package math takes some precaculus and automates it
package math

// simplify reevaluate, simplify, reevaluate

// Term the interface for part of an expression
// expression is going to be a composition of terms
type Term interface {
	// if the var is alone then the nubmber is zero
	// every term has a coefficietn even if that coeficient is one
	SetNumericCoefficient(int) Term
	SetVariableCoefficient(rune) Term // x or y or empty '' rune, convert any to lowercase
	// each side of an equation is an expression
	// types of coeficeint positive, negative, fractional
	SetIsSquareRooted(bool) // bool

	SetExponent(int) Term
	GetTerm() Term
	// is it a paranthesized expression, is it a square rooted expression
	// is it a paranthesized term, is it a square rooted term
	// use slices to hold expressions so that size can be easily determined
	// use slices to hold terms.
	// perhaps i can do this with the same interface types , since there will be different types of terms
}

// type Exponent struct {
// 	// if there is no exponent set this value to 1
// 	Ex int
// }
type TManufacturingDirector struct {
	builder Term
}

var tmaninstance *TManufacturingDirector

// term product

func GetManInstance() *TManufacturingDirector {
	if taninstance == nil {
		maninstance = new(TManufacturingDirector)
	}
	return tmaninstance
}

// TConstruct construct method
func (f *TManufacturingDirector) TConstruct(termcount int, abssets int, operators int) {
	f.builder.SetNumericCoefficient(NumericCoefficient).SetVariableCoefficient(VariableCoefficient).IsSquareRooted(bool).SetExponent(Exponent)
}

//EManufacturingDirector Expression Manufacturing Director
func (f *EManufacturingDirector) SetTBuilder(t Term) {
	t.builder = b
}

type TermP struct {
	NumericCoefficient  int
	VariableCoefficient rune
	IsSquareRooted      bool
	Exponent            int
}

// type TermBuilder struct {
// 	// math term
// 	mt TermProduct
// }

type ConstantTern struct {
	v TermP
}

func (t *TermP) SetNumericCoefficient(n int) *TermP {
	t.Number = n
	return t
}
func (t *TermP) SetVariableCoefficient(v rune) *TermP {
	t.VariableCoefficient = v
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
