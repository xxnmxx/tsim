package tsim

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	env := NewEnviroment()
	input := `new c;
	new a;
	a;
	crt a.rv = rev:100:v10`
	l := NewLexer(input)
	p := NewParser(l)
	program := p.ParseProgram()
	evaluated := Eval(program, env)

	tt := []struct {
		e ObjectType
	}{
		{e: CORP_OBJ},
		{e: CORP_OBJ},
		{e: CORP_OBJ},
		{e: CORP_OBJ},
	}
	for _, stmt := range program.Statements {
		fmt.Printf("len:%v\ttype:%T\n", len(program.Statements), stmt)
	}
	for i, test := range tt {
		if evaluated.Type() != test.e {
			t.Errorf("i:%v,e:%v,a:%v\n%+v", i, test.e, evaluated.Type(), evaluated)
		}
	}
}
