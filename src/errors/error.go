package errors

import (
	"ares/src/token"
	"fmt"
	"os"
)

const (
	SyntaxError = "Syntax Error"
)

type ErrorType string

type Error struct {
	Type ErrorType
	Msg  string
	Loc  token.Loc
}

func (e Error) Print() {
	fmt.Println(fmt.Sprintf("%v: %v\n\tat line %v, column %v", e.Type, e.Msg, e.Loc.Row, e.Loc.Col))
	os.Exit(1)
}
