package tsim

import "testing"

func TestReadChar(t *testing.T) {
	input := `crt=.1`
	l := NewLexer(input)
	tt := []struct{
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
		l.readChar()
	}
}
