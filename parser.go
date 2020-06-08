package tsim

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}


func (p *Parser) ParseProgram() *Program {
	program := &Program
	program.Statements = []Statement{}

	for curToken != EOF {
		stmt := parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statement, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case NEW:
		return p.parseNewStatement()
	default:
		return nil
	}
}

func (p *Parser) parseNewStatement() *NewStatement {
	stmt := NewStatement{Token: p.Token}
	if !p.expectPeek(IDENT) {
		return nil
	}
	stmt.Name = &Identifier{Token:p.curToken,Value:p.curToken.Literal}
	return stmt
}

// Helper functions
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return curToken == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return peekToken == t
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}
