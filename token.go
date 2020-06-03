package tsim

type TokenType string

const (
	// Special
	EOF = ""
	ILLEGAL = "Illegal"

	// Expression
	IDENT  = "Identifier"
	FLOAT  = "Float"
	ASSIGN = "="

	// Corp handling
	NEW = "New"

	// Acc handling
	CRT = "Create"
	RED = "Read"
	UPD = "Update"
	DEL = "Delete"

	// Sep
	DOT = "."
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(t TokenType, l string) Token {
	return Token{Type: t, Literal: l}
}

var keyword = map[string]Token{
	"new": Token{Type: NEW, Literal:"new"},
	"crt": Token{Type: CRT, Literal: "crt"},
	"red": Token{Type: RED, Literal: "red"},
	"upd": Token{Type: UPD, Literal: "upd"},
	"del": Token{Type: DEL, Literal: "del"},
}

func lookupKey(literal string) Token {
	tok, ok := keyword[literal]
	if !ok {
		return newToken(IDENT, literal)
	}
	return tok
}
