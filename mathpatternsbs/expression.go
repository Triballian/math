package math

// you are going to want the aperator and the location of that operator
type Expression interface {
	SetTermCount(int) Expression
	//absolute values
	SetABSsets(int) Expression
	SetOperators(int) Expression
	GetMathExpression() ExpressionProduct
}

// Director
type ManufacturingDirector struct {
	builder Expression
	count   int
}

//start singleton
//single Manufacturing director
// type singleManD struct {
// 	count int
// }

var maninstance *ManufacturingDirector

func GetManInstance() *ManufacturingDirector {
	if maninstance == nil {
		maninstance = new(ManufacturingDirector)
	}
	return maninstance
}

//for testing purposese only
func (m *ManufacturingDirector) AddOne() int {
	m.count++
	return m.count
}

//for testing purposese only
func (m *ManufacturingDirector) GetCount() int {
	return m.count
}

// construct method
func (f *ManufacturingDirector) Construct(termcount int, abssets int, operators int) {
	f.builder.SetTermCount(termcount).SetABSsets(abssets).SetOperators(operators)
}

//setBuilder Method
func (f *ManufacturingDirector) SetBuilder(b Expression) {
	f.builder = b
}

type ExpressionProduct struct {
	TermCount int
	ABSsets   int
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

func (m *MonomialBuilder) SetABSsets(abs int) Expression {
	m.v.ABSsets = abs
	return m
}

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
func (b *BinomialBuilder) SetABSsets(abs int) Expression {
	b.v.ABSsets = abs
	return b
}

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
func (t *TrinomialBuilder) SetABSsets(abs int) Expression {
	t.v.ABSsets = abs
	return t
}

func (t *TrinomialBuilder) SetOperators(o int) Expression {
	t.v.Operators = o
	return t
}

func (t *TrinomialBuilder) GetMathExpression() ExpressionProduct {
	return t.v
}
