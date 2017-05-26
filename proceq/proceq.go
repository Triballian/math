package proceq

import (
	"fmt"

	"github.com/Triballian/MATH/proceq/procterm"
)

func procEq(eq []string) {

	// eConsts := make([]string, 0, 10)

	pols := sepEx(eq) //polynomials

	hsfrst := hsFrst(pols)

	terms := sepTerms(eq)

	lopertrs := getLopertrs(hsfrst[0])

	fmt.Println("terms is:%s", terms[0])

	//proces the terms

	// eTerm := procterm.Term{}

	eterms = make([]procterm.Term, len(tem)s, 10)
	t *[]procterm.Term = &eterms

	position := 0

	for i, term := range terms {
		// r = strings.Split(term, `**`)
		// c := string.Split(term, x, y)
		// if len(r) < 2
		if yVar, present := procterm.TypeTerm(term, "y"); present {
			position++
			vterms := make(map[int][]procterm.Term)
			vterms = apend(vterms, vterms[int].position, position)
			// eTerm.position = position
			eterms.vTerm[i].Var = "y"
			eterms.vTerm[i].Coefficient = yVar[0]
			if eOp, present := procterm.TypTer(yVar[1], "**"); present {
				eterms.vTerm[i].Exponent = int(eOp[2])
			} else if _, present := procterm.TypeTerm(term, "x"); present {
				position++
				eterms.vTerm[i].Var = "x"
			} else {
				eterms.vTerm[i].Var = ""
				eTerm.ExConst.Number = term
			}

			// if xVar, present := procterm.TypeTerm(term, "x"); present {
			// 	position++
			// 	eTerm.position = position
			// 	eTerm.Xvar.XV = true
			// 	if eOp, present := procterm.TypTer(xVar[1], "**"); present {
			// 		eTerm.Xvar.ExponX.exp = int(eOp[2])
			// 	}
		}
		// does this term have a variable
		// inside every if set a value to indecate what it is instde expression struct
		// does this term have an exponent

	}

}
