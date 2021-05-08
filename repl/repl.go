package repl

import (
	"bufio"
	"fmt"
	"ginter/lexer"
	"ginter/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		fmt.Printf(PROMPT)
		line := scanner.Text()
		if line != "" {
			l := lexer.New(line)

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		}
	}
}
