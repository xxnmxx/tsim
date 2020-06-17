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

var accmap = map[string]AccType{
	ASSET:     Asset,
	LIABILITY: Liability,
	REVENUE:   Revenue,
	EXPENCE:   Expence,
}

func LookupAccToken(literal string) AccType {
	at, ok := accmap[literal]
	if ok {
		return at
	}
	return ""
}

var vatmap = map[string]VatType{
	v10:  V10,
	v8:   V8,
	v8r:  V8R,
	v5:   V5,
	v0:   V0,
	vf:   VF,
	vn:   VN,
	a10:  A10,
	a10t: A10t,
	a10c: A10c,
	a10n: A10n,
	a8:   A8,
	a8t:  A8t,
	a8c:  A8c,
	a8n:  A8n,
	a8r:  A8R,
	a8rt: A8Rt,
	a8rc: A8Rc,
	a8rn: A8Rn,
	a5:   A5,
	a5t:  A5t,
	a5c:  A5c,
	a5n:  A5n,
	af:   AF,
	an:   AN,
}


func LookupVatToken(literal string) VatType {
	vt, ok := vatmap[literal]
	if ok {
		return vt
	}
	return ""
}
