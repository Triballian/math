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

func prmptforInput(d int) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Example1 with one value in the Range: -5 7 Example2 with 3 values in the Range: 4 2;6;-2")
	fmt.Println("Please enter the first domain %d value and also the range value it is assigned to separated by semicolons then press Enter, type go and enter when you are done:", d + 1)

	fmt.Print("\n>>")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return text
}

// var f = new([]values.Fset)


func main() {
	d := 0
	var fsets = make(map[int]values.Fset)

	// v := values.Fset

	for {

		text := prmptforInput(d + 1)
		proctext(fsets, d)
		if text == "go" {
			fmt.Println("text is %s", text)
			fmt.Println("The results are.")
			break
		}
		
	}

}
