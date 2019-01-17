package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/zeroFruit/zf-lang/object"

	"github.com/zeroFruit/zf-lang/vm"

	"github.com/zeroFruit/zf-lang/compiler"

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

	constants := []object.Object{}
	globals := make([]object.Object, vm.GlobalsSize)
	symbolTable := compiler.NewSymbolTable()

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

		comp := compiler.NewWithState(symbolTable, constants)
		err := comp.Compile(prog)
		if err != nil {
			fmt.Fprintf(out, "Compilation failed:\n %s\n", err)
			continue
		}

		code := comp.Bytecode()
		constants = code.Constants

		machine := vm.NewWithGlobalsStore(code, globals)
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.LastPoppedStackElem()

		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "  parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
