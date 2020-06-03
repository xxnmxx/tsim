package tsim

type TokenType string

const (
	// Special
	EOF = ""

	// Expression
	IDENT  = "Identifier"
	FLOAT  = "Float"
	ASSIGN = "="

	// Acc handling
	CRT = "Create"
	RED = "Read"
	UPD = "Update"
	DEL = "Delete"

	// Infix
	DOT = "."
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(t TokenType, l string) Token {
	return &Token{Type: t, Literal: l}
}

// WIP
func lookupKey() Token {
	keyword := map[string]Token{
		"crt": Token{Type: CRT, Literal: "crt"},
		"red": Token{Type: RED, Literal: "red"},
		"upd": Token{Type: UPD, Literal: "upd"},
		"del": Token{Type: DEL, Literal: "del"},
	}

}
