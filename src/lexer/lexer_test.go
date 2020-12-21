package lexer

import (
	"fmt"
	"heron/src/token"
	"testing"
)

func TestScanner(t *testing.T) {
	input := `
	

	selector:hover {
    	background-color: blue; //ignore this
	} //ignore this as well`

	//TODO: add support for rule value that include whitespace
	l := New(input)
	//l.readChar()

	expected := []token.Token{
		{Literal: "selector", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "hover", Type: token.IDENT},
		{Literal: "{", Type: token.LBRACE},

		{Literal: "background-color", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "blue", Type: token.IDENT},
		{Literal: ";", Type: token.EOS},

		{Literal: "}", Type: token.RBRACE},
	}

	for _, expectedTok := range expected {
		tok := l.NextToken()

		fmt.Println(fmt.Sprintf("Type: %v; Literal: %v; Row: %v; Col: %v", tok.Type, tok.Literal, tok.Loc.Row, tok.Loc.Col))

		if tok.Type != expectedTok.Type {
			t.Errorf("Unexpected token type: expected %v instead of %v", expectedTok.Type, tok.Type)
		}

		if tok.Literal != expectedTok.Literal {
			t.Errorf("Unexpected token literal: expected %v instead of %v", expectedTok.Literal, tok.Literal)
		}
	}
}
