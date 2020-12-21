package errors

import (
	"heron/src/token"
	"testing"
)

func TestNewError(t *testing.T) {
	err := Error{Type: SyntaxError, Msg: "bla bla", Loc: token.Loc{Row: 10, Col: 5}}
	err.Print()
}
