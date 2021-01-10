package ast

import "heron/src/token"

type (
	//Program contains all the abstract syntax tree nodes
	Program struct {
		FileName string
		Imports  []Program
		Rules    []Selector
	}

	//Rule represents the key-value pair of a CSS property
	Rule struct {
		Name  string
		Value string
	}

	//Selector represents a CSS rule
	Selector struct {
		SelectorText string

		Nested []Selector
		Rules  []Rule

		Loc token.Loc
	}
)
