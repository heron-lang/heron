package scanner

import (
	"ares/src/token"
	"fmt"
	"testing"
)

func TestScanner(t *testing.T) {
	input := `
	selector {
		color: red;
		background-color: blue;
	}`

	l := New(input)
	l.readChar()

	expected := []token.Token{
		{Literal: "selector", Type: token.IDENT},
		{Literal: "{", Type: token.LBRACE},
		{Literal: "color", Type: token.IDENT},
		{Literal: "red", Type: token.IDENT},
		{Literal: "background-color", Type: token.IDENT},
		{Literal: "blue", Type: token.IDENT},
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
