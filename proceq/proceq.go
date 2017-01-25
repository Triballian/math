package proceq

func procEq(eq string) {

	type exConst struct {
		number int
	}
	eConsts := make([]string, 0, 10)

	pols := sepEx(eq) //polynomials

	hsfrst := hsFrst(pols)

	terms := sepTerms(eq)

	lopertrs := getLopertrs(hsfrst[0])

}
