package tsim

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}
