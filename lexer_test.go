package tsim

import "testing"

func TestNextToken(t *testing.T) {
	input := 'crt=.'
	l := NewLexer(input)
	tt := {
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
			t.Errorf("i: %v,e: %v, a: %v",i,test.e,string(l.ch))
		}
	}
}
