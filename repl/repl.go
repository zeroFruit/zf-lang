package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/zeroFruit/zf-lang/parser"

	"github.com/zeroFruit/zf-lang/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	defer func() {
		recover()
		Start(in, out)
	}()

	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		prog := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}

		io.WriteString(out, prog.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "  parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
