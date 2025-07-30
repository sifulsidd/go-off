package repl

import (
	"bufio"
	"go/scanner"

	"fmt"

	"io"

	"go-work/lexer"

	"go-work/token"
)

const PROMPT = ">> "

// takes an input reader and output writer 
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// this is similar to a while loop, basically just keeps our program running 
	for {
		// prints out prompt first  
		fmt.Fprintf(out, PROMPT)

		// basically reads the next line of input
		scanned := scanner.Scan()
		
		// if it is empty jsut return nothing
		if !scanned {
			return
		}

		// otherwise gets the text we received
		line := scanner.Text()

		// basically creates a new lexer with the line
		l := lexer.New(line)
		
		// go through each character of the lexer and print out the tokens as long as it isn't end of range
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// similar to the f"{variable}" in python where the +v is the variable in the string, and the tok is the string it will be 
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

