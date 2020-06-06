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
	upd c.exp = 2.00`
	l := NewLexer(input)
	tt := []struct {
		eType    TokenType
		eLiteral string
	}{
		{eType: NEW, eLiteral: "new"},
		{eType: IDENT, eLiteral: "a"},
		{eType: SEMICOLON, eLiteral: ";"},
		{eType: CREATE, eLiteral: "crt"},
		{eType: IDENT, eLiteral: "c"},
		{eType: DOT, eLiteral: "."},
		{eType: IDENT, eLiteral: "rev"},
		{eType: ASSIGN, eLiteral: "="},
		{eType: FLOAT, eLiteral: "100"},
		{eType: UPDATE, eLiteral: "upd"},
		{eType: IDENT, eLiteral: "c"},
		{eType: DOT, eLiteral: "."},
		{eType: IDENT, eLiteral: "exp"},
		{eType: ASSIGN, eLiteral: "="},
		{eType: FLOAT, eLiteral: "2.00"},
	}
	for i, test := range tt {
	tok := l.NextToken()
		if tok.Type != test.eType {
			t.Errorf("typeError\ni: %v\tetype: %v\tatype: %v\n", i, test.eType, tok.Type)
		}
		if tok.Literal != test.eLiteral {
			t.Errorf("literalError\ni: %v\teliteral: %v\taliteral: %v", i, test.eLiteral, tok.Literal)
		}
	}
}
