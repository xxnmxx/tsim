package tsim

import "testing"

func TestReadChar(t *testing.T) {
	input := `crt=.1`
	l := NewLexer(input)
	tt := []struct {
		e string
	}{
		{"c"},
		{"r"},
		{"t"},
		{"="},
		{"."},
	}
	for i, test := range tt {
		if string(l.ch) != test.e {
			t.Errorf("i: %v,e: %v, a: %v", i, test.e, string(l.ch))
		}
		l.readChar()
	}
}

func TestNextToken(t *testing.T) {
	input := `new a;
	crt c.rev = 100
	upd c.exp = 1.00`
	l := NewLexer(input)
	tt := []struct {
		eType    TokenType
		eLiteral string
	}{
		{eType: NEW, eLiteral: "new"},
		{eType: IDENT, eLiteral: "c"},
		{eType: SEMICOLON, eLiteral: ";"},
		{eType: CRT, eLiteral: "crt"},
		{eType: IDENT, eLiteral: "c"},
		{eType: DOT, eLiteral: "."},
		{eType: IDENT, eLiteral: "rev"},
		{eType: ASSIGN, eLiteral: "="},
		{eType: FLOAT, eLiteral: "100"},
		{eType: UPD, eLiteral: "upd"},
		{eType: IDENT, eLiteral: "c"},
		{eType: DOT, eLiteral: "."},
		{eType: IDENT, eLiteral: "exp"},
		{eType: ASSIGN, eLiteral: "="},
		{eType: FLOAT, eLiteral: "1.00"},
	}
	tok := l.NextToken()
	for i, test := range tt {
		if tok.Type == test.eType && tok.Literal == test.eLiteral {
			t.Errorf("i: %v,etype: %v,eliteral %v, atype: %v, aliteral: %v", i, test.eType, test.eLiteral, tok.Type, tok.Literal)
		}
		l.NextToken()
	}
}
