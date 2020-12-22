package lexer

import (
	"heron/src/errors"
	"heron/src/token"
	"strings"
)

//Lexer is a structure that represents the scanner
type Lexer struct {
	input []byte

	position     int
	nextPosition int
	loc          token.Loc

	ch byte
}

//New creates new lexer from input
func New(input []byte) *Lexer {
	l := &Lexer{input: input, loc: token.Loc{Row: 1, Col: 0}}
	l.readChar()
	return l
}

//NextToken creates new token and advances characters
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '/':
		l.readChar()

		switch l.ch {
		case '/':
			for l.ch != '\n' {
				l.readChar()
			}

			return l.NextToken()
		case '*':
			l.readChar() //skip asterisk
			for l.ch != '*' && l.input[l.nextPosition] != '/' {
				l.readChar()
			}

			l.readChar() //skip asterisk
			l.readChar() //skip backslash

			return l.NextToken()
		default:
			err := &errors.Error{Msg: "illegal character", Type: errors.SyntaxError, Loc: l.loc}
			err.Print()
		}
	case '{':
		tok = l.newToken(token.LBRACE)
	case '}':
		tok = l.newToken(token.RBRACE)
	case ':':
		tok = l.newToken(token.COLON)
	case ';':
		tok = l.newToken(token.EOS)
	//case '\n':
	//	tok = l.newToken(token.EOS)
	case 0:
		tok.Literal = "EOF"
		tok.Type = token.EOF
		return tok
	default:
		tok.Type = token.IDENT
		tok.Loc = l.loc
		tok.Literal = l.eat(l.isIdent)
		return tok //prevents program from reading next char
	}

	l.readChar()
	return tok
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

	if l.ch == '\n' {
		l.loc.Col = 0
		l.loc.Row++
	}
}

func (l *Lexer) eat(check func() bool) string {
	var eaten strings.Builder

	for check() {
		eaten.WriteByte(l.ch)
		l.readChar()
	}

	return eaten.String()
}

func (l *Lexer) newToken(tt token.TokenType) token.Token {
	return token.Token{Type: tt, Literal: string(l.ch), Loc: l.loc}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) isIdent() bool {
	return l.isLetter() || l.isNumber() || l.ch == '-' || l.ch == '*' || l.ch == '#' || l.ch == '_' || l.ch == '%'
}

func (l *Lexer) isLetter() bool {
	return 'a' <= l.ch && l.ch <= 'z' || 'A' <= l.ch && l.ch <= 'Z'
}

func (l Lexer) isNumber() bool {
	return ('0' <= l.ch && l.ch <= '9') || l.ch == '.'
}
