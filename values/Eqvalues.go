// storage for both each row of domain with its range
// to be compared to see if the equation is a function

package values

type fdomain struct {
	domain int
}

type dset struct {
	Domset []fdomain
}

type frange struct {
	Range []int
}

// Fmap function map of the Damain to the range.
type Fmap struct {
	dset
	frange
}
