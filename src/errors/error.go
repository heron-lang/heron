package errors

import (
	"fmt"
	"heron/src/token"
	"os"
)

const (
	//SyntaxError is used when the Heron code does not match our syntax
	SyntaxError = "SyntaxError"

	//ImportError is thrown when the compiler fails to fetch the specified file
	ImportError = "ImportError"
)

//ErrorType represents the type of error
type ErrorType string

//Error is an organized and consistent representation of an error message
type Error struct {
	//Type is the type of error
	Type ErrorType

	//Msg is the error message
	Msg string

	//Loc is the location of where the error originated
	Loc token.Loc
}

//Print will format and log the error to the console
func (e Error) Print() {
	fmt.Println(fmt.Sprintf("%v: %v\n\tat line %v, column %v", e.Type, e.Msg, e.Loc.Row, e.Loc.Col))
	os.Exit(1)
}
