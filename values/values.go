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

// Fmap function map of the Damain to the range.
// fap map should be a collection of domain and range instances
type Fmap struct {
	Fdomain
	Frange
}
