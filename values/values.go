// storage for both each row of domain with its range
// to be compared to see if the equation is a function

package values

type Fdomain struct {
	Domain int
}

// type Dset struct {
// 	Domset []fdomain
// }

type Frange struct {
	Range []int
}

// Fsets function map of the Damain to the range.
// fsets should be a collection of domain and range instances
type Fset struct {
	Fdomain
	Frange
}


