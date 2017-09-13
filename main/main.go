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

func proctext(t string, d int) {
	// dnr := []string
	// r := []string
	// fd = values.Fdomain
	// fr = values.Frange
	fm = values.Fmap
	dnr := strings.Fields(t)
	r := strings.Replace(dnr[1], ";", ",", -1)

	// fd, _ = strconv.Atoi(dnr[0])
	fm.domain, _ = strconv.Atoi(dnr[0])


	for _, i := range r {
		j, err := strconv.Atoi(fmt.Sprint("s", i))
		if err != nil {
			panic(err)
		}
		fm.Range = append(fm.Range, j)
	}
	f = append(fm.domain, fm.range)

}

func prmptforInput(d int) string {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Example1 with one value in the Range: -5 7 Example2 with 3 values in the Range: 4 2;6;-2")
	fmt.Println("Please enter the first domain %d value and also the range value it is assigned to separated by semicolons then press Enter, type results and enter when you are done:", d)

	fmt.Print("\n>>")

	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return text
}

var f = new([]values.Fmap)

func main() {
	d := 0

	// v := values.Fmap

	for {
		text := prmptforInput(d + 1)
		if text == "quit" {
			// fmt.Println("text is %s", text)
			fmt.Println("The results are.")
			break
		}
		proctext(text, d)
	}

}
