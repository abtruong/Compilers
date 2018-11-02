package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	var cfgStack = stack.New() // Empty stack
	var accept bool            // Accept state
	var currStack string       // Contents of stack (to keep track/display)
	var buf bytes.Buffer       // Buffer for concatenating currStack string efficiently

	inputStrings := [3]string{ // Input strings
		"(a+a)*a$",
		"a*(a/a)$",
		"a(a+a)$",
	}

	fmt.Println("\n\n************************************")
	fmt.Println("*     PREDICTIVE PARSING TABLE     *")
	fmt.Print("************************************\n\n")

	for _, inputString := range inputStrings { // Begin to check if input string will be accepted or rejected
		fmt.Print(" ----------------------------------\n\n")
		fmt.Print(" *****\tTRACING ", inputString, "\n\n")

		index := 0 // Set index to 0 (start of input string, ie: "(" )
		read := string(inputString[index])
		fmt.Print("\tREAD:\t ", read, "\n\n")

		cfgStack.Push("$") // Push $ (bottom of stack)
		fmt.Println("\tPUSHED:\t", "$")
		currStack = "$"
		fmt.Print(" *****\tSTACK: \t ", currStack, "\n\n")

		cfgStack.Push("E") // Push E (first CFG)
		fmt.Println("\tPUSHED:\t", "E")
		currStack = "$E"
		fmt.Print(" *****\tSTACK: \t ", currStack, "\n\n")

		for true {
			popped := cfgStack.Pop().(string)                 // Pop top of stack  $E --> $
			currStack = strings.TrimSuffix(currStack, popped) // Remove the stack's top element from current stack elements
			fmt.Print("\tPOPPED:\t ", popped, "\n *****\tSTACK: \t ", currStack, "\n")

			if currStack != "" { // Buffer to write new content of stack after pushing new elements (if pushed)
				buf.WriteString(currStack)
			}

			for popped == "1" { // 1 = lambda; while lambda is top-most on stack, pop lambda off and pop again
				popped = cfgStack.Pop().(string)
				currStack = strings.TrimSuffix(currStack, popped)
				fmt.Print("\tPOPPED:\t ", popped, "\n *****\tSTACK: \t ", currStack, "\n")
				buf.Reset()
				buf.WriteString(currStack)
			}

			if popped == "$" && read == popped { // Stack is now empty and
				fmt.Print("\tMATCH: \t ", popped, "\n")
				currStack = strings.TrimSuffix(currStack, popped)
				fmt.Print("\tPOPPED:\t ", popped, "\n *****\tSTACK: \t", currStack, "\n\n")
				accept = true
			}

			if accept == true {
				fmt.Print(" *****\t", inputString, " IS ACCEPTED.\n\n")
				break
			}

			for popped == read { // If popped element is the same as the input string's read, then pop again, assign read to next input string element
				fmt.Println("\tMATCH: \t", popped)
				popped = cfgStack.Pop().(string) // $QT( --> $QT (second pop after read)
				fmt.Print("\tPOPPED:\t ", popped, "\n")
				index++

				read = string(inputString[index]) // _(_ a+a)*a$ --> ( _a_ +a)*a$
				fmt.Println("\tREAD:  \t", read)

				currStack = strings.TrimSuffix(currStack, popped) // $QT( --> $QT
				fmt.Print(" *****\tSTACK: \t ", currStack, "\n")
				buf.Reset()
				buf.WriteString(currStack)
			}

			grammar := string(trace(popped, read)) // Traces the grammar to the correct table grammar (trace.go file)
			fmt.Print("\tTRACE: \t [", popped, ",", read, "] = ", grammar, "\n\n")

			if grammar == "NULL" { // Tracing led to empty state; reject and move to next input string
				fmt.Print(" *****\t", inputString, " IS REJECTED.\n\n")
				break
			}

			for i := len(grammar) - 1; i >= 0; i-- { // Push the contents of the CFG after tracing
				buf.WriteString(string(grammar[i]))
				cfgStack.Push(string(grammar[i]))
				fmt.Print("\tPUSHED:\t ", string(grammar[i]), "\n")
			}

			currStack = buf.String()
			buf.Reset()
			fmt.Print(" *****\tSTACK: \t ", currStack, "\n\n")
		}

		buf.Reset()
		accept = false // Reset accept state for next input string
	}
}
