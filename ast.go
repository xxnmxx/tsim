package tsim

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// Statements
type NewStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ns *NewStatement) TokenLiteral() string { return ns.Token.Literal }
func (ns *NewStatement) statementNode()       {}

// Expressions
type Identifier struct {
	Token Token // IDENT
	Value string
}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode()      {}

type CorpLiteral struct {
	Token Token // CORP
	Body  *Corp
}

func (cp *Corp) TokenLiteral() string { return cp.Token.Literal }
func (cp *Corp) expressionNode()      {}
