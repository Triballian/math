package procterm

// type ExConst struct { // expression constant
// 	Term
// 	number int
// }

// type ECoefficient struct { // Expression
// 	Term
// 	number int
// }

// type XVar struct { // expression with X Variable Term
// 	Term
// 	// EVar string
// 	XV bool
// }

// type YVar struct { // expression with y Variable Term
// 	Term
// 	// EVar string
// 	XV bool
// }

// type ExponX struct { // X variable with exponent
// 	XVar
// 	exp int
// }

// type ExponY struct { // y variable with
// 	YVar
// 	exp int
// }

// type ExponXCoef struct { // x variable with exponent and coefficient
// 	Term
// 	ExponX
// 	ECoefficient
// }

// type ExponYCoef struct { // y variable with exponent and coefficient
// 	Term
// 	ExponY
// 	ECoefficient
// }

type Term struct {
	Position int // needing to deferintiate terms is why position here is important
	Expression
	Var        string // x or y
	Coefficiet int
	Exponent   int
	Number     int
}

// type Sterm struct {
// 	Position int
// 	Expression
// }

// type Tterm struct {
// 	Position int
// 	Expression
// }

type Expression struct {
	postion int
}
