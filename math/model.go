// Package math takes some precaculus and automates it
package math

// simplify reevaluate, simplify, reevaluate

// Term the interface for part of an expression
// expression is going to be a composition of terms
type Term interface {
	// if the var is alone then the nubmber is zero
	// every term has a coefficietn even if that coeficient is one
	SetIsABSed(bool) Term
	SetIsParenthesesed(bool) Term
	SetIsSquareRooted(bool) Term // bool
	SetExponent(int) Term
	SetNumericCoefficient(int) Term
	SetVariableCoefficient(rune) Term // x or y or empty '' rune, convert any to lowercase
	// each side of an equation is an expression
	// types of coeficeint positive, negative, fractional

	GetTerm() TermP
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

func GetTManInstance() *TManufacturingDirector {
	if tmaninstance == nil {
		tmaninstance = new(TManufacturingDirector)
	}
	return tmaninstance
}

// TConstruct construct method
func (f *TManufacturingDirector) TConstruct(isABSed bool, isParenthesesed bool, isSquareRooted bool, exponent int, numericCoefficient int, variableCoefficient rune) {
	f.builder.SetIsABSed(isABSed).SetIsParenthesesed(isParenthesesed).SetIsSquareRooted(isSquareRooted).SetExponent(exponent)
	f.builder.SetNumericCoefficient(numericCoefficient).SetVariableCoefficient(variableCoefficient)
}

//EManufacturingDirector Expression Manufacturing Director
func (f *TManufacturingDirector) SetTBuilder(t Term) {
	f.builder = t
}

type TermP struct {
	IsABSed             bool
	IsParenthesesed     bool
	SquareRooted        bool
	Exponent            int
	NumericCoefficient  int
	VariableCoefficient rune
}

// type TermBuilder struct {
// 	// math term
// 	mt TermProduct
// }
type ABSedTerm struct {
	v TermP
}

func (c *ABSedTerm) SetIsABSed(s bool) Term {
	c.v.IsABSed = s
	return c
}
func (c *ABSedTerm) SetIsParenthesesed(s bool) Term {
	c.v.IsParenthesesed = s
	return c
}
func (c *ABSedTerm) SetIsSquareRooted(s bool) Term {
	c.v.SquareRooted = s
	return c
}
func (c *ABSedTerm) SetExponent(e int) Term {
	c.v.Exponent = e
	return c
}
func (c *ABSedTerm) SetNumericCoefficient(n int) Term {
	c.v.NumericCoefficient = n
	return c
}
func (c *ABSedTerm) SetVariableCoefficient(v rune) Term {
	c.v.VariableCoefficient = v
	return c
}

func (c *ABSedTerm) GetTerm() TermP {
	return c.v
}

type ConstantTerm struct {
	v TermP
}

//even thought the types are created virtually identically we can identify instance for its specific type symplifying or even dictating the operations that are applied to it.
func (c *ConstantTerm) SetIsABSed(s bool) Term {
	c.v.IsABSed = s
	return c
}
func (c *ConstantTerm) SetIsParenthesesed(s bool) Term {
	c.v.IsParenthesesed = s
	return c
}
func (c *ConstantTerm) SetIsSquareRooted(s bool) Term {
	c.v.SquareRooted = s
	return c
}
func (c *ConstantTerm) SetExponent(e int) Term {
	c.v.Exponent = e
	return c
}
func (c *ConstantTerm) SetNumericCoefficient(n int) Term {
	c.v.NumericCoefficient = n
	return c
}
func (c *ConstantTerm) SetVariableCoefficient(v rune) Term {
	c.v.VariableCoefficient = v
	return c
}

func (c *ConstantTerm) SetIsABSed(s bool) Term {
	c.v.IsABSed = s
	return c
}
func (c *ConstantTerm) SetIsParenthesesed(s bool) Term {
	c.v.IsParenthesesed = s
	return c
}
func (c *ConstantTerm) SetIsSquareRooted(s bool) Term {
	c.v.SquareRooted = s
	return c
}
func (c *ConstantTerm) SetExponent(e int) Term {
	c.v.Exponent = e
	return c
}
func (c *ConstantTerm) SetNumericCoefficient(n int) Term {
	c.v.NumericCoefficient = n
	return c
}
func (c *ConstantTerm) SetVariableCoefficient(v rune) Term {
	c.v.VariableCoefficient = v
	return c
}

// GetTerm get term function
// func (t *Term) Setvars(thisvar string, varval bool) *TermP {
// 	if thisvar == "x" || thisvar == "X" {
// 		t.Xvar = varval
// 	}
// 	if thisvar == "y" || thisvar == "Y" {
// 		t.Yvar = varval
// 	}

// 	return t
// }

func (c *ConstantTerm) GetTerm() TermP {
	return c.v
}

type VarTerm struct {
	v TermP
}

func (c *VarTerm) SetIsABSed(s bool) Term {
	c.v.IsABSed = s
	return c
}
func (c *VarTerm) SetIsParenthesesed(s bool) Term {
	c.v.IsParenthesesed = s
	return c
}
func (c *VarTerm) SetIsSquareRooted(s bool) Term {
	c.v.SquareRooted = s
	return c
}
func (c *VarTerm) SetExponent(e int) Term {
	c.v.Exponent = e
	return c
}
func (c *VarTerm) SetNumericCoefficient(n int) Term {
	c.v.NumericCoefficient = n
	return c
}
func (c *VarTerm) SetVariableCoefficient(v rune) Term {
	c.v.VariableCoefficient = v
	return c
}

func (c *VarTerm) GetTerm() TermP {
	return c.v
}

type VarCoeefficientTerm struct {
	v TermP
}

func (c *VarCoeefficientTerm) SetIsABSed(s bool) Term {
	c.v.IsABSed = s
	return c
}
func (c *VarCoeefficientTerm) SetIsParenthesesed(s bool) Term {
	c.v.IsParenthesesed = s
	return c
}
func (c *VarCoeefficientTerm) SetIsSquareRooted(s bool) Term {
	c.v.SquareRooted = s
	return c
}

func (c *VarCoeefficientTerm) SetExponent(e int) Term {
	c.v.Exponent = e
	return c
}
func (c *VarCoeefficientTerm) SetNumericCoefficient(n int) Term {
	c.v.NumericCoefficient = n
	return c
}
func (c *VarCoeefficientTerm) SetVariableCoefficient(v rune) Term {
	c.v.VariableCoefficient = v
	return c
}

func (c *VarCoeefficientTerm) GetTerm() TermP {
	return c.v
}

type SquareRootedTerm struct {
	v TermP
}

func (c *SquareRootedTerm) SetIsABSed(s bool) Term {
	c.v.IsABSed = s
	return c
}
func (c *SquareRootedTerm) SetIsParenthesesed(s bool) Term {
	c.v.IsParenthesesed = s
	return c
}
func (c *SquareRootedTerm) SetIsSquareRooted(s bool) Term {
	c.v.SquareRooted = s
	return c
}
func (c *SquareRootedTerm) SetExponent(e int) Term {
	c.v.Exponent = e
	return c
}
func (c *SquareRootedTerm) SetNumericCoefficient(n int) Term {
	c.v.NumericCoefficient = n
	return c
}
func (c *SquareRootedTerm) SetVariableCoefficient(v rune) Term {
	c.v.VariableCoefficient = v
	return c
}

func (c *SquareRootedTerm) GetTerm() TermP {
	return c.v
}

type ExponentedTerm struct {
	v TermP
}

func (c *ExponentedTerm) SetIsABSed(s bool) Term {
	c.v.IsABSed = s
	return c
}
func (c *ExponentedTerm) SetIsParenthesesed(s bool) Term {
	c.v.IsParenthesesed = s
	return c
}
func (c *ExponentedTerm) SetExponent(e int) Term {
	c.v.Exponent = e
	return c
}
func (c *ExponentedTerm) SetIsSquareRooted(s bool) Term {
	c.v.SquareRooted = s
	return c
}
func (c *ExponentedTerm) SetNumericCoefficient(n int) Term {
	c.v.NumericCoefficient = n
	return c
}
func (c *ExponentedTerm) SetVariableCoefficient(v rune) Term {
	c.v.VariableCoefficient = v
	return c
}

func (c *ExponentedTerm) GetTerm() TermP {
	return c.v
}
