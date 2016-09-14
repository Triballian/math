package main

// determine if the equation is a function
// for ever input x there is only one output y
// solve for the output y
// term1 + or minus term2 ... = y
import (
	//"bufio"
	"bufio"
	"fmt"
	"os"
	//"io"
	"log"
	//"strings"
	"unicode"
	"strings"
	//"strconv"

)

func main() {


	// primaryoperators := []string{`+`, `-`}
	//infinite loop to continue evaluating the input expressions

	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nType the equation to be identified as either a function or not a function")
		fmt.Print("\n>>")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		var ftext string

		ftext = strings.Map(func(t rune) rune {
			if unicode.IsSpace(t) {
				return -1
			}
			return t
		}, text)

		expression := strings.Split(ftext, `=`)

		if text == "quit\n" {
			fmt.Println("Exited at user request.")
			break
		}
		if len(expression[0]) < len(expression[1]) {
			fmt.Printf("The answer portions:%s. The math portion:%s\n", expression[0], expression[1])
		} else {
			fmt.Printf("The answer portions:%s. The math portion:%s\n", expression[1], expression[0])
		}

		fmt.Printf("left side of equation:%s, Right side of equtioin:%s\n", expression[0], expression[1])
		// need a slice of number left to right for expression zero
		// need a slice of operators from left to right for expression0
		bufferstrng := make([]string, 0, 10)
		exzeronumbers := make([]string, 0, 10)
		exzerooperators := make([]string, 0, 10)
		var bs string
		counter := 0



		for _, value := range expression[0] {
			counter++

			//s, _ := strconv.Atoi(value)

			s := fmt.Sprintf("%c", value)


			if len(bufferstrng) == 0 {
				//s = fmt.Sprintf("%c", value)

				bufferstrng = append(bufferstrng, s)

				continue
			}
			if s == `+` || s == `-`{
				fmt.Printf("bufferstring is %s", bufferstrng)
				bs = strings.Join(bufferstrng, " ")
				fmt.Printf("bs is %s \n", bs)
				//i, _ = strconv.ParseInt(bs, 10, 64)
				exzeronumbers = append(exzeronumbers, bs)
				exzerooperators = append(exzerooperators, s)
				bufferstrng = bufferstrng[:0]
				continue
			}
			bufferstrng = append(bufferstrng, s)




			fmt.Printf("The value is:%c \n", value)
			fmt.Printf("counter is %s \n", counter)
			if counter == len(expression[0]){
				bs = strings.Join(bufferstrng, " ")
				exzeronumbers = append(exzeronumbers, bs)
			}




		}
		fmt.Printf("exzeronumbers :%s \n", exzeronumbers)
		fmt.Printf("exzerooperators :%s \n", exzerooperators)

	}
}
