package scanner

import "ares/src/token"

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '{':
		tok = l.newToken(token.LBRACE)
	case '}':
		tok = l.newToken(token.RBRACE)
	case ':':
		tok = l.newToken(token.COLON)
	case ';':
		tok = l.newToken(token.SEMICOLON)
	default:
		if l.isLetter() {
			tok.Type = token.IDENT
			tok.Literal = l.eat(l.isLetter)
			return tok
		} else if l.isDigit() {
			tok.Type = token.INT
			tok.Literal = l.eat(l.isDigit)
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}

	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) eat(check func() bool) string {
	start := l.position
	for check() {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) newToken(tt token.TokenType) token.Token {
	return token.Token{Type: tt, Literal: string(l.ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) isLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || 'A' <= l.ch && l.ch <= 'Z' || l.ch == '_'
}

func (l *Lexer) isDigit() bool {
	return '0' <= l.ch && l.ch <= '9'
}
