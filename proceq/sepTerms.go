package proceq

import "strings"

func sepTerms(lsEx string) []string {
	// bufferSlice := make([]string, 0, 10)
	// e := make([]string, 0, 10)
	// opexoperators := make([]string, 0, 10)
	// var bs string
	// counter := 0

	// for _, value := range lsEx {
	//     counter++

	//     //s, _ := strconv.Atoi(value)

	//     //s := fmt.Sprintf("%c", value)
	f := func(c rune) bool {
		return c == '+' || c == '-'
	}

	bufferSlice := strings.FieldsFunc(lsEx, f)

	// if value == `+` || value == `-` {
	//     fmt.Printf("bufferstring is %s \n", bufferslice)
	//     bs = strings.Join(bufferslice, "")
	//     fmt.Printf("bs is %s \n", bs)
	//     //i, _ = strconv.ParseInt(bs, 10, 64)
	//     e = append(e, bs)
	//     opexoperators = append(opexoperators, value)
	//     bufferslice = bufferslice[:0]
	//     continue
	// }
	//     bufferslice = append(bufferslice, value)

	//     fmt.Printf("The value is:%s \n", value)
	//     fmt.Printf("counter is %d \n", counter)
	//     if counter == len(ex[0]) {
	//         fmt.Println(counter)
	//         bs = strings.Join(bufferslice, "")
	//         e = append(e, bs)
	//     }

	// }
	// fmt.Printf("opexoperators :%s \n", opexoperators)
	return bufferSlice
}
