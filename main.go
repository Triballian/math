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
	//"strings"
)

//numofterms := 2
//term := []string{2x**2, 3y**2}
//func getterm(s string) string {
//
//}
//for i := 0; i < range term; i++{
//}
//for i := 0; i < numofterms; i++ {
//fmt.Println("%s", term1)
//}
func main() {
	//infinite loop to continue evaluating the input expressions
	// var in string
	// var fmtdequation string = ""
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Type the equation to be identified as either a function or not a function")
		fmt.Print("\n>>")
		//line := in.ReadString(" ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
		if text == "quit\n" {
			break
		}

		//remove spaces -1 teslls Replace that there is no limit to the number of replacements
		//fmtdequation := strings.Replace(equation, " ", "", -1)

		//fmt.Println(fmtdequation)
		//fmt.Println(equation)
		// sepreate by =
		// seperate by arithic operator

		//for i := 0; i < numofterms; i++ {
		//	fmt.Println("%s", term1)
	}
}

// fmt.Println(fmtdequation)
//}

//func simpleReadLine(r io.Reader) (string, error) {
//	rl := bufio.NewScanner(r)
//	rl.Scan()
//	if err := rl.Err(); err != nil {
//		return "", err
//	}
//	return rl.Text(), nil
//}
//
//func main() {
//	r := strings.NewReader(fmt.scanln())
//	fmt.Println(simpleReadLine(r))
//}
