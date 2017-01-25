package proceq

func hsFrst(s []string) []string {

	if len(s[0]) < len(s[1]) {
		s[1], s[0] = s[0], s[1]
	}
	return s
}
