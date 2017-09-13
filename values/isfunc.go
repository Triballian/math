package values

// Isfunc returns a bool based on weather or not the expression is a function
func (f Fmap) Isfunc() bool {
	if len(f.Range) > 1 {
		return false
	}
	return true
}
