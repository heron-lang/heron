package lexer

import (
	"ares/src/token"
	"fmt"
	"testing"
)

func TestScanner(t *testing.T) {
	input := `
	selector:hover {
		color: red;
		background-color: blue;
	}`

	//TODO: add support for rule value that include whitespace
	l := New(input)
	l.readChar()

	expected := []token.Token{
		{Literal: "selector", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "hover", Type: token.IDENT},
		{Literal: "{", Type: token.LBRACE},

		{Literal: "color", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "red", Type: token.IDENT},
		{Literal: ";", Type: token.SEMICOLON},

		{Literal: "background-color", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "blue", Type: token.IDENT},
		{Literal: ";", Type: token.SEMICOLON},
		{Literal: "}", Type: token.RBRACE},
	}

	for _, expectedTok := range expected {
		tok := l.NextToken()
		fmt.Println(tok)
		if tok.Type != expectedTok.Type {
			t.Errorf("Unexpected token type: expected %v instead of %v", expectedTok.Type, tok.Type)
		}

		if tok.Literal != expectedTok.Literal {
			t.Errorf("Unexpected token literal: expected %v instead of %v", expectedTok.Literal, tok.Literal)
		}
	}
}
