package procterm

type exConst struct {   // expression constant
    term
	number int
}

type eCoefficient struct {     // Expression 
    term
    number int
}

type xVar struct {               // expression with X Variable Term
    term
    eVar string := x
}

type yVar struct{               // expression with y Variable Term
    term
    eVar string := y
}

type exponX struct {                // X variable with exponent
    xVar
    exp     int
}

type exponY struct {                // y variable with 
    yVar
    exp     int
}

type exponXCoef struct {            // x variable with exponent and coefficient
    term
    exponX
    eCoefficient
}

type exponYCoef  struct {            // y variable with exponent and coefficient
    term
    exponY
    eCoefficient
}

type term struct{
    position int
    expression
}

type expression struct{ 
    Eside   string
    size    int         // number of terms
}


