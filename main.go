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
	"log"
	"strings"
	"unicode"
)


func main() {
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
		fmt.Printf("left side of equation:%s, Right side of equtioin:%s\n", expression[0], expression[1])


	}
}


