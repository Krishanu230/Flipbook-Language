package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Krishanu230/Flipbook-Language/evaluator"
	"github.com/Krishanu230/Flipbook-Language/lexer"
	"github.com/Krishanu230/Flipbook-Language/object"
	"github.com/Krishanu230/Flipbook-Language/parser"
	"github.com/Krishanu230/Flipbook-Language/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "  parser errors:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
