package tsim

type TokenType string

const (
	// Special
	EOF     = ""
	ILLEGAL = "ILLEGAL"
	// Expression
	IDENT  = "Identifier"
	FLOAT  = "FLOAT"
	ASSIGN = "="
	// Keywords
	NEW = "NEW"
	CRT = "CRT"
	RED = "RED"
	UPD = "UPD"
	DEL = "DEL"
	// Sep
	DOT = "."
	SEMICOLON = ";"
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(t TokenType, l string) Token {
	return Token{Type: t, Literal: l}
}

var keyword = map[string]Token{
	"new": Token{Type: NEW, Literal: "new"},
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
