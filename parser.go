package tsim

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(l) *Parser {
}
