package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	IDENT = "IDENT"
	INT   = "INT"

	//Delimiters
	SEMICOLON = ";"
	COLON     = ":"

	LBRACE = "{"
	RBRACE = "}"

	ILLEGAL = "ILLEGAL"
)
