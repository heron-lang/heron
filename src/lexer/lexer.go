package lexer

import "heron/src/token"

type Lexer struct {
	input string

	position     int
	nextPosition int
	loc          token.Loc

	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input, loc: token.Loc{Row: 1, Col: 0}}
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
		tok = l.newToken(token.EOS)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		tok.Type = token.IDENT
		tok.Loc = l.loc
		tok.Literal = l.eat(l.isIdent)
		return tok //prevents program from reading next char
	}

	l.readChar()
	return tok
}

func (l *Lexer) isIdent() bool {
	return l.isLetter() || l.ch == '-' || l.ch == '*' || l.ch == '.' || l.ch == '#'
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.loc.Col++
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
	return token.Token{Type: tt, Literal: string(l.ch), Loc: l.loc}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()

		if l.ch == '\n' {
			l.loc.Col = 0
			l.loc.Row++
		}
	}
}

func (l *Lexer) isLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || 'A' <= l.ch && l.ch <= 'Z' || l.ch == '_'
}
