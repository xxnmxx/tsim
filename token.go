package tsim

type TokenType string

const (
	// Special
	EOF     = ""
	ILLEGAL = "ILLEGAL"
	// Expression
	IDENT  = "IDENT"
	FLOAT  = "FLOAT"
	ASSIGN = "="
	CORP = "CORP"
	// Keywords
	NEW    = "NEW"
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
	// Sep
	DOT       = "."
	SEMICOLON = ";"
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(t TokenType, ch byte) Token {
	return Token{Type: t, Literal: string(ch)}
}

var keyword = map[string]TokenType{
	"new": NEW,
	"crt": CREATE,
	"rd": READ,
	"upd": UPDATE,
	"del": DELETE,
}

func LookupKey(literal string) TokenType {
	tok, ok := keyword[literal]
	if ok {
		return tok
	}
	return IDENT
}
