package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Loc     Loc
}

type Loc struct {
	Row int
	Col int
}

const (
	IDENT = "identifier"

	//Delimiters
	EOS       = "EOS"
	SEMICOLON = ";"
	COLON     = ":"

	LBRACE = "{"
	RBRACE = "}"

	EOF = "EOF"
)
