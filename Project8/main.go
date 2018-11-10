package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	var currStack string
	var buf bytes.Buffer

	inputStrings := [2]string {
		"(i+i)*i$",
		"(i*)$",
	}

	fmt.Println("\n\n************************************")
	fmt.Println("*     PREDICTIVE PARSING TABLE     *")
	fmt.Print("************************************\n\n")

	for _, inputString := range inputStrings {
		fmt.Print(" ----------------------------------\n\n")
		fmt.Print(" *****\tTRACING: ", inputString, "\n\n")

		var inputStack = stack.New()
		index := 0

		inputStack.Push("0")
		currStack = "0 "
		fmt.Println("\tPUSHED:\t", "0")

		read := string(inputString[index])
		fmt.Print("\tREAD:\t ", read, "\n\n")

		fmt.Print(" *****\tSTACK:\t ", currStack, "\n\n")
		
		for true {
			popped := inputStack.Pop().(string)
			currStack = strings.TrimSuffix(currStack, popped + " ")
			fmt.Print("\tPOPPED:\t ", popped, "\n *****\tSTACK:\t ", currStack, "\n")
			buf.WriteString(currStack)

			parse := trace(popped, read) // parse = table[row][col] 
			fmt.Print("\tTRACE:\t [", popped, ",", read, "] = ", parse, "\n")

			if strings.ToUpper(parse) == "ACC" {
				buf.Reset()
				fmt.Print("\n *****\t", inputString, " IS ACCEPTED.\n\n")
				break;
			} else if strings.ToUpper(parse) == "" {
				buf.Reset()
				fmt.Print("\n *****\t", inputString, " IS REJECTED.\n\n")
				break;
			}

			char := string(parse[0])

			if strings.ToUpper(char) == "S" {
				inputStack.Push(popped)
				buf.WriteString(popped + " ")
				fmt.Println("\tPUSHED:\t", popped)

				inputStack.Push(read)
				buf.WriteString(read + " ")
				fmt.Println("\tPUSHED:\t", read)
				
				restOfParse := strings.TrimPrefix(parse, char)
				inputStack.Push(restOfParse)
				buf.WriteString(restOfParse + " ")
				fmt.Println("\tPUSHED:\t", restOfParse)

				index++
				read = string(inputString[index])
				fmt.Println("\tREAD:\t", read)

			} else if strings.ToUpper(char) == "R" {
				inputStack.Push(popped)
				buf.Reset()
				buf.WriteString(currStack)
				buf.WriteString(popped + " ")
				currStack = buf.String()
				fmt.Println("\tPUSHED:\t", popped)
				fmt.Println(" *****\tSTACK:\t", currStack)

				popAmount, nonTerminal := getRuleLengthAndNonTerm(parse)
				fmt.Println(" *****\tPOPPING:", popAmount, "ELEMENTS OFF STACK:")

				for i := 0; i < popAmount; i++ {
					popped = inputStack.Pop().(string)
					fmt.Println("\tPOPPED:\t", popped)
					currStack = strings.TrimSuffix(currStack, popped + " ")
				}
				fmt.Println(" *****\tSTACK:\t", currStack)

				popped = inputStack.Pop().(string)
				currStack = strings.TrimSuffix(currStack, popped + " ")
				fmt.Println("\tPOPPED:\t", popped, "\n *****\tSTACK:\t", currStack)
				parse := trace(popped, nonTerminal)
				fmt.Print("\tTRACE:\t [", popped, ",", nonTerminal, "] = ", parse, "\n")
				buf.Reset()
				buf.WriteString(currStack)

				inputStack.Push(popped)
				buf.WriteString(popped + " ")
				fmt.Println("\tPUSHED:\t", popped)

				inputStack.Push(nonTerminal)
				buf.WriteString(nonTerminal + " ")
				fmt.Println("\tPUSHED:\t", nonTerminal)

				inputStack.Push(parse)
				buf.WriteString(parse + " ")
				fmt.Println("\tPUSHED:\t", parse)
			}

			currStack = buf.String()
			fmt.Print(" *****\tSTACK:\t ", currStack, "\n\n")
			buf.Reset()
		}
	}
}
