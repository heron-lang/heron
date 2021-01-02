package lexer

import (
	"fmt"
	"heron/src/token"
	"testing"
)

func TestComments(t *testing.T) {
	input := `
	*/
		IGNORE ME!
	*/

	//ignore this as well`

	//TODO: add support for rule value that include whitespace

	createTest(t, input, []token.Token{})
}

func TestEOS(t *testing.T) {
	input := `
		selector:hover {
			font: 100%
		}	
	`

	expected := []token.Token{
		{Literal: "selector", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "hover", Type: token.IDENT},
		{Literal: "{", Type: token.LBRACE},

		{Literal: "font", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "100%", Type: token.IDENT},

		{Literal: "}", Type: token.RBRACE},
	}

	createTest(t, input, expected)
}

func TestNumbers(t *testing.T) {
	input := `
		selector:hover {
			transition: color 1s;
		}	
	`

	expected := []token.Token{
		{Literal: "selector", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "hover", Type: token.IDENT},
		{Literal: "{", Type: token.LBRACE},

		{Literal: "transition", Type: token.IDENT},
		{Literal: ":", Type: token.COLON},
		{Literal: "color", Type: token.IDENT},
		{Literal: "1s", Type: token.IDENT},
		{Literal: ";", Type: token.EOS},

		{Literal: "}", Type: token.RBRACE},
	}

	createTest(t, input, expected)
}

func createTest(t *testing.T, input string, expected []token.Token) {
	l := New([]byte(input))

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
