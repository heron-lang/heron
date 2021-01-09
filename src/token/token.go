package token

//TokenType represents the type of token
type TokenType string

//Token represents a token
type Token struct {
	Type    TokenType
	Literal string
	Loc     Loc
}

//Loc represents a location in the program
type Loc struct {
	Row int
	Col int
}

const (
	//IDENT is short for identifier. It names the languages entities, this could be a property value or name.
	IDENT = "identifier"

	//STRING is a literal that contains the text between two quotation marks
	STRING = "string"

	//EOS is an end-of-statement. It is typically represented as a semicolon
	EOS = "EOS"

	//COLON is a delimiter is used for separating rule properties
	COLON = ":"

	//LBRACE is a delimiter that signifies the start of a rule
	LBRACE = "{"

	//RBRACE is a delimiter that signifies the end of a rule
	RBRACE = "}"

	//EOF signifies the end of a file
	EOF = "EOF"

	//ATRULE is the prefix that starts commands like "import"
	ATRULE = "@"
)
