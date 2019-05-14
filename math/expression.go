package math

// you are going to want the aperator and the location of that operator
type Expression interface {
	SetTermCount(int) Expression
	//absolute values
	// SetABSsets(int) Expression
	SetOperators(int) Expression
	GetMathExpression() ExpressionProduct
}

// EManufacturingDirector Manufacturing Director
type EManufacturingDirector struct {
	builder Expression
}

var maninstance *EManufacturingDirector

func GetManInstance() *EManufacturingDirector {
	if maninstance == nil {
		maninstance = new(EManufacturingDirector)
	}
	return maninstance
}

// construct method
func (f *EManufacturingDirector) Construct(termcount int, operators int) {
	f.builder.SetTermCount(termcount).SetOperators(operators)
}

//EManufacturingDirector Expression Manufacturing Director
func (f *EManufacturingDirector) SetBuilder(b Expression) {
	f.builder = b
}

type ExpressionProduct struct {
	TermCount int
	// ABSsets   int
	Operators int
}

// //expression builder
// type ExBuilder struct {
// 		// math term
// 		mt TermProduct
// 	}

type MonomialBuilder struct {
	v ExpressionProduct
}

func (m *MonomialBuilder) SetTermCount(c int) Expression {
	if c == 1 {
		m.v.TermCount = c

	}
	return m
}

// func (m *MonomialBuilder) SetABSsets(abs int) Expression {
// 	m.v.ABSsets = abs
// 	return m
// }

func (m *MonomialBuilder) SetOperators(o int) Expression {
	m.v.Operators = o
	return m
}

func (m *MonomialBuilder) GetMathExpression() ExpressionProduct {
	return m.v
}

type BinomialBuilder struct {
	v ExpressionProduct
}

func (b *BinomialBuilder) SetTermCount(c int) Expression {
	if c == 2 {
		b.v.TermCount = c

	}
	return b
}

// func (b *BinomialBuilder) SetABSsets(abs int) Expression {
// 	b.v.ABSsets = abs
// 	return b
// }

func (b *BinomialBuilder) SetOperators(o int) Expression {
	b.v.Operators = o
	return b
}

func (b *BinomialBuilder) GetMathExpression() ExpressionProduct {
	return b.v
}

type TrinomialBuilder struct {
	v ExpressionProduct
}

func (t *TrinomialBuilder) SetTermCount(c int) Expression {
	if c == 3 {
		t.v.TermCount = c

	}
	return t
}

// func (t *TrinomialBuilder) SetABSsets(abs int) Expression {
// 	t.v.ABSsets = abs
// 	return t
// }

func (t *TrinomialBuilder) SetOperators(o int) Expression {
	t.v.Operators = o
	return t
}

func (t *TrinomialBuilder) GetMathExpression() ExpressionProduct {
	return t.v
}
