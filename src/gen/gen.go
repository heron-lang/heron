package gen

import (
	"ares/src/ast"
)

type Gen struct {
	curNode ast.Selector
}

func (g *Gen) Generate(program *ast.Program) string {
	var css string

	for _, selector := range program.Rules {
		g.curNode = selector

		css += g.genRules()
	}

	return css
}

func (g Gen) genRules() (css string) {
	css += g.curNode.SelectorText + "{"

	for _, rule := range g.curNode.Rules {
		css += rule.Name + ":" + rule.Value + ";"
	}

	css += "}"
	return
}
