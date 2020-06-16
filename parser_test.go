package tsim

import (
	"fmt"
	"testing"
)

func TestParseProgram(t *testing.T) {
	input := `new c;
	crt c.sales = rev:100.0:v10;
	crt c.costs = exp:222:a10;
	c;`
	l := NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	for i, statement := range program.Statements {
		fmt.Printf("%v\t%+v\n",i,statement)
	}
	tt := []struct {
		e string
	}{
		{e: "*tsim.NewStatement"},
		{e: "*tsim.CreateStatement"},
		{e: "*tsim.CreateStatement"},
		{e: "*tsim.ExpressionStatement"},
	}

	for i, test := range tt {
		a := fmt.Sprintf("%T", program.Statements[i])
		if a != test.e {
			t.Errorf("i: %v, e: %T, a:%T\n", i, test.e, program.Statements[i])
		}
	}
}
