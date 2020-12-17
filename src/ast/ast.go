package ast

type Program struct {
	Rules []Selector
}

type Rule struct {
	Name  string
	Value string
}

type Selector struct {
	SelectorText string
	Rules        []Rule
	Loc          Loc
}

type Loc struct {
	Row int
	Col int
}
