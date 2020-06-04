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
	}
	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition++
}

// ToDo
func (l *Lexer) NextToken() Token {
	l.eraseSpace()
	switch l.ch {
	case '.':
		return newToken(DOT, string(l.ch))
	case '=':
		return newToken(ASSIGN, string(l.ch))
	case 0:
		return newToken(EOF, string(0))
	default:
		if l.isLetter(l.ch) {
			literal := l.readIdentifier()
			return lookupKey(literal)
		}
		if l.isDigit(l.ch) {
			preLiteral := l.readDigit()
			surLiteral := ""
			if l.ch == '.' {
				surLiteral = l.readDigit()
			} else {
				return newToken(FLOAT, surLiteral)
			}
			literal := preLiteral + "." + surLiteral
			return newToken(FLOAT, literal)
		}
		return newToken(ILLEGAL, ILLEGAL)
	}
}

func (l *Lexer) peekChar() byte {
	return l.input[l.readPosition]
}

func (l *Lexer) isLetter(c byte) bool {
	return 'a' <= c && 'z' <= c || 'A' <= c && 'Z' <= c || c == '_'
}

func (l *Lexer) isDigit(c byte) bool {
	return '0' <= c && '9' <= c
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) readDigit() string {
	position := l.position
	for l.isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

func (l *Lexer) eraseSpace() {
	if l.ch == ' ' {
		l.readChar()
	}
}
