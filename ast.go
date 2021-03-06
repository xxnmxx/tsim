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
	Value *Corp
}

func (ns *NewStatement) statementNode()       {}
func (ns *NewStatement) TokenLiteral() string { return ns.Token.Literal }

type CreateStatement struct {
	Token Token // CREATE token
	Attr  *Identifier
	Name  *Identifier
	Value *AccLiteral
}

func (cs *CreateStatement) statementNode()       {}
func (cs *CreateStatement) TokenLiteral() string { return cs.Token.Literal }

type ExpressionStatement struct {
	Token Token // First token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {return es.Token.Literal}
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

func (cp *CorpLiteral) TokenLiteral() string { return cp.Token.Literal }
func (cp *CorpLiteral) expressionNode()      {}

type AccLiteral struct {
	Token    Token // ACC
	AccToken Token // REVENUE etc
	Value    Token // FLOAT
	VatToken Token // V8 etc
}

func (al *AccLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *AccLiteral) expressionNode()      {}
