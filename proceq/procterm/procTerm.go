package procterm

func ProcTerm(terms []string, t *[]Term) *[]Term {
	position := 0
	for i, term := range terms {
		// r = strings.Split(term, `**`)
		// c := string.Split(term, x, y)
		// if len(r) < 2
		if yVar, present := TypeTerm(term, "y"); present {
			position++
			eterm := Term{}
			*t[i] = eterm
			// eTerm.position = position
			eterms.vTerm[i].Var = "y"
			eterms.vTerm[i].Coefficient = yVar[0]
			if eOp, present := TypTerm(yVar[1], "**"); present {
				eterms.vTerm[i].Exponent = int(eOp[2])
			} else if _, present := TypeTerm(term, "x"); present {
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
