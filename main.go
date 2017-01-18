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

func prmptforInput() string {
		
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nType the equation to be identified as either a function or not a function")
		fmt.Print("\n>>")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)			
		}
		return strings.TrimSpace(text)
}

func sepOperandNanswer(t string) ([]string, []string, []string){
		f := make([]string, 0, 20)
		var fstring string

		for _, value := range t {
			if unicode.IsSpace(value) {
				continue
			}
			fstring = fmt.Sprintf("%c", value)
			f = append(f, fstring)
		}
		fstring = strings.Join(f, "")

		e := strings.Split(fstring, `=`)
		// fmt.Println("expression = %s", expression)

		a := make([]string, 0, 10)
		o := make([]string, 0, 10)
		if len(e[0]) < len(e[1]) {
			fmt.Printf("The answer portions:%s. The math portion:%s\n", e[0], e[1])
			a = mkslice(e[0])
			o = mkslice(e[1])
		} else {
			fmt.Printf("The answer portions:%s. The math portion:%s\n", e[1], e[0])
			a = mkslice(e[1])
			o = mkslice(e[0])
		}
		
		return a, o, e
}

func getexopnumbers(o []string, ex []string) []string {
		bufferstrng := make([]string, 0, 10)
		e := make([]string, 0, 10)
		opexoperators := make([]string, 0, 10)
		var bs string
		counter := 0

		for _, value := range o {
			counter++

			//s, _ := strconv.Atoi(value)

			//s := fmt.Sprintf("%c", value)

			if value == `+` || value == `-` {
				fmt.Printf("bufferstring is %s \n", bufferstrng)
				bs = strings.Join(bufferstrng, "")
				fmt.Printf("bs is %s \n", bs)
				//i, _ = strconv.ParseInt(bs, 10, 64)
				e = append(e, bs)
				opexoperators = append(opexoperators, value)
				bufferstrng = bufferstrng[:0]
				continue
			}
			bufferstrng = append(bufferstrng, value)

			fmt.Printf("The value is:%s \n", value)
			fmt.Printf("counter is %d \n", counter)
			if counter == len(ex[0]) {
				fmt.Println(counter)
				bs = strings.Join(bufferstrng, "")
				e = append(e, bs)
			}

		}
		fmt.Printf("opexoperators :%s \n", opexoperators)
		return e
}

func procexop(e []string) (int, int){
		var divisor int
		var exsuffix int
		lenexopno := len(e)
		if lenexopno > 1 {
			for i, _ := range e {
				sn := 0

				en := 0
				exchar := 0

				if lenexopno > 1 {

					for _, value := range e[i] {
						lenexopno = len(e[i])

						if fmt.Sprintf("%c", value) == "x" {
							fmt.Printf("number %s, when x = 0 is 0.\n", e[i])
							e[i] = "0"

							en = 0
							break

						} else if fmt.Sprintf("%c", value) == "y" {
							sn = en + 2
							divisor, _ = strconv.Atoi(fmt.Sprintf("%s", e[i][:en]))
							//divisor, _ = strconv.Atoi(e[i][:en])
							//fmt.Printf("else if divisor: %s en:%d\n e[:0]:%s \n", divisor, en, e[i][:1])
							en++
							continue

						} else if fmt.Sprintf("%c", value) == "*" {
							exchar++

						} else if fmt.Sprintf("%c", value) != "*" && exchar == 2 {
							sn++
							en++

							if en == lenexopno {

								exsuffix, _ = strconv.Atoi(e[i][sn:en])
								fmt.Printf("exsuffix :%s , en:%d, sn:%d\n", e[i][sn:en], en, sn)
							}
							continue
						}

						en++
					}
				}
			}
		}
		return divisor, exsuffix
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
		
		if text == "quit" {
			// fmt.Println("text is %s", text)
			fmt.Println("Exited at user request.")
			break
		}
		//end commands
		

		
		
		// opex answerex start
		answerex, opex, expression := sepOperandNanswer(text)
		fmt.Printf("opex:%s, answerex:%s\n", opex, answerex)
		// opex answerex end
		// need a slice of number left to right for expression zero
		// need a slice of operators from left to right for expression0
		// get exopnumbers begin
		exopnumbers := getexopnumbers(opex, expression)
		fmt.Printf("exopnumbers :%s \n", exopnumbers)
		
		// get exopnumbers end
		// procexop start
		divisor, exsuffix := procexop(exopnumbers)
		// procexop end
		// x is the domain y is the range, there can be only one y answer as a fuction of x.
		fmt.Printf("exopnumbers :%s, exsuffix :%d, divisor :%d \n", exopnumbers, exsuffix, divisor)
		fmt.Printf("y**%d = %s/%d\n", exsuffix, answerex, divisor)
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
