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
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case NEW:
		return p.parseNewStatement()
	case CREATE:
		return p.parseCreateStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseNewStatement() *NewStatement {
	stmt := &NewStatement{Token: p.curToken}
	if !p.expectPeek(IDENT) {
		return nil
	}
	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}
	stmt.Value = NewCorp()
	p.nextToken()
	if p.peekTokenIs(SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseCreateStatement() *CreateStatement {
	stmt := &CreateStatement{Token: p.curToken}
	if !p.expectPeek(IDENT) {
		return nil
	}
	stmt.Name = &Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	if !p.expectPeek(DOT) {
		return nil
	}
	if !p.expectPeek(IDENT) {
		return nil
	}
	stmt.Attr = &Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
	if !p.expectPeek(ASSIGN) {
		return nil
	}
	if !p.expectPeek(REVENUE) && !p.expectPeek(EXPENCE) && !p.expectPeek(ASSET) && !p.expectPeek(LIABILITY) {
		return nil
	}
	al := p.parseAccLiteral()
	stmt.Value = al
	p.nextToken()
	return stmt
}

func (p *Parser) parseAccLiteral() *AccLiteral {
	al := AccLiteral{}
	al.Token = Token{Type: ACC, Literal: "ACC"}
	al.AccToken = p.curToken
	p.nextToken()
	al.Value = p.curToken
	p.nextToken()
	al.VatToken = p.curToken
	return &al
}

// Need extend
func (p *Parser) parseExpressionStatement() *ExpressionStatement {
	stmt := &ExpressionStatement{Token:p.curToken}
	stmt.Expression = p.parseIdentifier()
	p.nextToken()
	if p.peekTokenIs(SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseIdentifier() Expression {
	return &Identifier{Token:p.curToken,Value:p.curToken.Literal}
}

// Helper functions
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	return false
}
