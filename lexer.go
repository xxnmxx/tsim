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
	l.ch = input[readPosition]
	l.position = l.readPosition
	l.readPosition++
}

// ToDo
func (l *Lexer) NextToken() Token {
	switch l.ch {
	case '=':
		return newToken(ASSIGN,string(l.ch))
	default:
		if isLetter(l.ch) {
		}
		if isDigit(l.ch) {
		}
	}
}

func (l *Lexer) isLetter(c byte) bool {
	return 'a' <= c && 'z' <= c || 'A' <= c && 'Z' <= c || c == '_'
}

func (l *Lexer) isDigit(c byte) bool {
	return '0' <= '9'
}
