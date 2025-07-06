package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hertzcodes/compiler-in-go/src/interpreter/token"

	"github.com/hertzcodes/compiler-in-go/src/interpreter/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		line := scanner.Scan()
		if !line {
			return
		}
		text := scanner.Text()
		l := lexer.New(text)
		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Printf("%+v\n", t)
		}
	}
}
