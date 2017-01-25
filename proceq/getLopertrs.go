package proceq

func getLopertrs(lsEx string) []string {
	o := make([]string, 0, 5)
	for _, value := range lsEx {
		if value == '+' {
			o = append(o, "+")
		}
		if value == '-' {
			o = append(o, "-")
		}

	}
	return o
}
