package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Triballian/math/values"
)

func proctext(fsets []values.Fset, d int) {
	// dnr := []string
	// r := []string
	// fd = values.Fdomain
	// fr = values.Frange
	// fm = values.Fsets
	dnr := strings.Fields(t)
	r := strings.Replace(dnr[1], ";", ",", -1)

	// fd, _ = strconv.Atoi(dnr[0])
	[d]fsets.Domain, _ = strconv.Atoi(dnr[0])



	for _, i := range r {
		j, err := strconv.Atoi(fmt.Sprint("s", i))
		if err != nil {
			panic(err)
		}
		[d]fsets.Range = append([d]ftests.Range, j)
	}
	f = append(fm.domain, [d]fsets.Range)

}