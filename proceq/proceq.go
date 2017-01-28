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
	var eterms []*procterm.Term{}

	term := make(map[int][]*Term)

	position := 0

	for i, term := range terms {
		// r = strings.Split(term, `**`)
		// c := string.Split(term, x, y)
		// if len(r) < 2
		if yVar, present := procterm.TypeTerm(term, "y"); present {
			position++
			eTerm.position = position
			eTerm.Yvar.YV = true
			eTerm.yvar.ECoefficient = yvar[0]
			if eOp, present := procterm.TypTer(yVar[1], "**"); present {
				eTerm.Yvar.ExponY.exp = int(eOp[2])
			} else if _, present := procterm.TypeTerm(term, "x"); present {
				position++
				eTerm.Xvar.XV = true
			} else {
				eTerm.ExConst.number = term
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
