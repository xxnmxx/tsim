package tsim

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// ToDo
func (l *Lexer) NextToken() Token {
	var tok Token
	l.eraseSpace()
	switch l.ch {
	case '.':
		tok = newToken(DOT, l.ch)
	case ';':
		tok = newToken(SEMICOLON, l.ch)
	case '=':
		tok = newToken(ASSIGN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if l.isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupKey(tok.Literal)
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
		if l.isDigit(l.ch) {
			preLiteral := l.readDigit()
			surLiteral := ""
			if l.ch == '.' {
				l.readChar()
				surLiteral = l.readDigit()
				literal := preLiteral + "." + surLiteral
				tok.Literal = literal
				tok.Type = FLOAT
				return tok
			} else {
				tok.Literal = preLiteral
				tok.Type = FLOAT
				return tok
			}
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	return l.input[l.readPosition]
}

func (l *Lexer) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readDigit() string {
	position := l.position
	for l.isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) eraseSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
