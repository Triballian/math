package main
// y=3x+1 is a funcion of x since for each real number x, the expression 3x + 1 is a unique real number.  
// determine if the equation is a function
// for ever input x there is only one output y
// solve for the output y
// term1 + or minus term2 ... = y
import (
	"bufio"
	"fmt"
	"os"
	//"io"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func prmptforInput() {
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nType the equation to be identified as either a function or not a function")
		fmt.Print("\n>>")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			break
			fmt.println("fatal error")
		}
		return text


}

func mkslice(s string) []string {

	slength := len(s)
	ss := make([]string, 0, slength)
	for _, v := range s {
		ss = append(ss, fmt.Sprintf("%c", v))
	}

	return ss

}

func main() {

	// primaryoperators := []string{`+`, `-`}
	//infinite loop to continue evaluating the input expressions

	for {
		//start input ui
		text := prmptforInput()

		//text = getuserinput
		
		// is commands? if yes return break need to create a map of commands. and isequation
		// todo creat http ui
		// use unit test
		//end input ui

		//start commands
		

		if text == "quit\n" {
			fmt.Println("Exited at user request.")
			break
		}
		//end commands

		
		ftext := make([]string, 0, 10)
		var fstring string

		for _, value := range text {
			if unicode.IsSpace(value) {
				continue
			}
			fstring = fmt.Sprintf("%c", value)
			ftext = append(ftext, fstring)
		}
		fstring = strings.Join(ftext, "")

		expression := strings.Split(fstring, `=`)
		//fmt.Println(expression)

		answerex := make([]string, 0, 10)
		opex := make([]string, 0, 10)
		if len(expression[0]) < len(expression[1]) {
			fmt.Printf("The answer portions:%s. The math portion:%s\n", expression[0], expression[1])
			answerex = mkslice(expression[0])
			opex = mkslice(expression[1])
		} else {
			fmt.Printf("The answer portions:%s. The math portion:%s\n", expression[1], expression[0])
			answerex = mkslice(expression[1])
			opex = mkslice(expression[0])
		}

		fmt.Printf("opex:%s, answerex:%s\n", opex, answerex)
		// need a slice of number left to right for expression zero
		// need a slice of operators from left to right for expression0
		bufferstrng := make([]string, 0, 10)
		exopnumbers := make([]string, 0, 10)
		opexoperators := make([]string, 0, 10)
		var bs string
		counter := 0

		for _, value := range opex {
			counter++

			//s, _ := strconv.Atoi(value)

			//s := fmt.Sprintf("%c", value)

			if value == `+` || value == `-` {
				fmt.Printf("bufferstring is %s \n", bufferstrng)
				bs = strings.Join(bufferstrng, "")
				fmt.Printf("bs is %s \n", bs)
				//i, _ = strconv.ParseInt(bs, 10, 64)
				exopnumbers = append(exopnumbers, bs)
				opexoperators = append(opexoperators, value)
				bufferstrng = bufferstrng[:0]
				continue
			}
			bufferstrng = append(bufferstrng, value)

			fmt.Printf("The value is:%c \n", value)
			fmt.Printf("counter is %d \n", counter)
			if counter == len(expression[0]) {
				fmt.Println(counter)
				bs = strings.Join(bufferstrng, "")
				exopnumbers = append(exopnumbers, bs)
			}

		}
		fmt.Printf("exopnumbers :%s \n", exopnumbers)
		fmt.Printf("opexoperators :%s \n", opexoperators)
		var divisor int
		var exsuffix int
		lenexopno := len(exopnumbers)
		if lenexopno > 1 {
			for i, _ := range exopnumbers {
				sn := 0

				en := 0
				exchar := 0

				if lenexopno > 1 {

					for _, value := range exopnumbers[i] {
						lenexopno = len(exopnumbers[i])

						if fmt.Sprintf("%c", value) == "x" {
							fmt.Printf("number %s, when x = 0 is 0.\n", exopnumbers[i])
							exopnumbers[i] = "0"

							en = 0
							break

						} else if fmt.Sprintf("%c", value) == "y" {
							sn = en + 2
							divisor, _ = strconv.Atoi(fmt.Sprintf("%s", exopnumbers[i][:en]))
							//divisor, _ = strconv.Atoi(exopnumbers[i][:en])
							//fmt.Printf("else if divisor: %s en:%d\n exopnumber[:0]:%s \n", divisor, en, exopnumbers[i][:1])
							en++
							continue

						} else if fmt.Sprintf("%c", value) == "*" {
							exchar++

						} else if fmt.Sprintf("%c", value) != "*" && exchar == 2 {
							sn++
							en++

							if en == lenexopno {

								exsuffix, _ = strconv.Atoi(exopnumbers[i][sn:en])
								fmt.Printf("exsuffix :%s , en:%d, sn:%d\n", exopnumbers[i][sn:en], en, sn)
							}
							continue
						}

						en++
					}
				}
			}
		}
		fmt.Printf("exopnumbers :%s, exsuffix :%d, divisor :%s \n", exopnumbers, exsuffix, divisor)
		fmt.Printf("y**%d = %d/%d\n", exsuffix, answerex, divisor)
		answer, _ := strconv.Atoi(strings.Join(answerex, ""))
		fmt.Printf("ano = %d / %d\n", answer, divisor)

		ano := answer / divisor
		fmt.Printf("y**%d = %d\n", exsuffix, ano)
		fmt.Printf("exsuffix:%d mod 2 = %d \n", exsuffix, exsuffix%2)
		if exsuffix%2 == 0 {
			fmt.Println("This is a function")
		}

	}
}

//func rmspace(fslice []string) []string{
//
//	for k, v := range fslice {
//		if unicode.IsSpace(v) {
//			continue
//		}
//		fslice = append(fslice[:k], fslice[k+1:]...)
//	}
//}
