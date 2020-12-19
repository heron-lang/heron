package gen

import (
	"ares/src/ast"
)

type Gen struct {
	Program *ast.Program
	Output  string
	curNode ast.Selector
}

func (g *Gen) Generate() {
	for _, selector := range g.Program.Rules {
		g.curNode = selector

		g.Output += g.genRules(g.curNode.SelectorText, g.curNode)
	}
}

func (g *Gen) genRules(selector string, node ast.Selector) (css string) {
	css += selector + "{"

	for _, rule := range node.Rules {
		css += rule.Name + ":" + rule.Value + ";"
	}

	for _, nested := range node.Nested {
		g.Output += g.genRules(selector+" "+nested.SelectorText, nested)
	}

	css += "}"
	return
}
