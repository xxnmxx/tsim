package tsim

import (
	"bufio"
	"io"
	"fmt"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := NewEnviroment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := NewLexer(line)
		p := NewParser(l)

		program := p.ParseProgram()
		evaluated := Eval(program,env)
		//fmt.Println(evaluated)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out,"\n")
		}
	}
}
