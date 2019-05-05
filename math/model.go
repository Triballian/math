package math

// math var expression var
type mxvar struct {
	xvar bool
}

type myvar struct {
	yvar bool
}

type mvar struct {
	mxvar
	myvar
}

type term struct {
	number int
	mvar   // check if one or the other is not nil
}
