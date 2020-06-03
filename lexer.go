package tsim

import "fmt"

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
	return '0' <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for l.isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.readPosition]
}

// Can not deal like 0.0.1111
func (l *Lexer) readFloat() (string, error) {
	position := l.position
	for l.isDigit(l.ch) {
		l.readChar()
	}
	if l.ch == '.' {
		for l.isDigit(l.ch) {
			l.readChar()
		}
		if l.ch == '.' {
			return "", fmt.Errorf("too many dots")
		}
		return l.input[position:l.readPosition], nil
	}
	return l.input[position:l.readPosition], nil
}
