package ast

import "heron/src/token"

//Program contains all the abstract syntax tree nodes
type Program struct {
	FileName string
	Imports  []Program
	Rules    []Selector
}

//Rule represents the key-value pair of a CSS property
type Rule struct {
	Name  string
	Value string
}

//Selector represents a CSS rule
type Selector struct {
	SelectorText string

	Nested []Selector
	Rules  []Rule

	Loc token.Loc
}
