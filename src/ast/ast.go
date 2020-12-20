package ast

import "heron/src/token"

type Program struct {
	Rules []Selector
}

type Rule struct {
	Name  string
	Value string
}

type Selector struct {
	SelectorText string

	Nested []Selector
	Rules  []Rule

	Loc token.Loc
}
