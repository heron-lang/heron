package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	IDENT = "IDENT"

	//Delimiters
	EOS       = "EOS"
	SEMICOLON = ";"
	COLON     = ":"

	LBRACE = "{"
	RBRACE = "}"

	EOF = "EOF"
)
