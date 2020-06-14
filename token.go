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
	CORP   = "CORP"
	// Keywords
	NEW    = "NEW"
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
	// Sep
	DOT       = "."
	SEMICOLON = ";"
	// Acc related
	ACC       = "ACC"
	ASSET     = "ASSET"
	LIABILITY = "LIABILITY"
	REVENUE   = "REVENUE"
	EXPENCE   = "EXPENCE"
	// VAT related
	v10  = "v10"
	v8   = "v8"
	v8r  = "v8r"
	v5   = "v5"
	v0   = "v0"
	vf   = "vf"
	vn   = "vn"
	a10  = "a10"
	a10t = "a10t"
	a10c = "a10c"
	a10n = "a10n"
	a8   = "a8"
	a8t  = "a8t"
	a8c  = "a8c"
	a8n  = "a8n"
	a8r  = "a8r"
	a8rt = "a8rt"
	a8rc = "a8rc"
	a8rn = "a8rn"
	a5   = "a5"
	a5t  = "a5t"
	a5c  = "a5c"
	a5n  = "a5n"
	af   = "af"
	an   = "an"
)

type Token struct {
	Type    TokenType
	Literal string
}

func newToken(t TokenType, ch byte) Token {
	return Token{Type: t, Literal: string(ch)}
}

var keyword = map[string]TokenType{
	// Literals
	"new": NEW,
	"crt": CREATE,
	"rd":  READ,
	"upd": UPDATE,
	"del": DELETE,
	// Acc related
	"ast": ASSET,
	"lbl": LIABILITY,
	"rev": REVENUE,
	"exp": EXPENCE,
	// Vat related
	"v10":  v10,
	"v8":   v8,
	"v8r":  v8r,
	"v5":   v5,
	"v0":   v0,
	"vf":   vf,
	"vn":   vn,
	"a10":  a10,
	"a10t": a10t,
	"a10c": a10c,
	"a10n": a10n,
	"a8":   a8,
	"a8t":  a8t,
	"a8c":  a8c,
	"a8n":  a8n,
	"a8r":  a8r,
	"a8rt": a8rt,
	"a8rc": a8rc,
	"a8rn": a8rn,
	"a5":   a5,
	"a5t":  a5t,
	"a5c":  a5c,
	"a5n":  a5n,
	"af":   af,
	"an":   an,
}

func LookupKey(literal string) TokenType {
	tok, ok := keyword[literal]
	if ok {
		return tok
	}
	return IDENT
}
