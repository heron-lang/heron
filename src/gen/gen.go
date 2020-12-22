package gen

import (
	"heron/src/ast"
	"strings"
)

type Gen struct {
	Program *ast.Program
	Output  strings.Builder
	curNode ast.Selector
}

func (g *Gen) Generate() {
	for _, selector := range g.Program.Rules {
		g.curNode = selector

		g.Output.WriteString(g.genRules(g.curNode.SelectorText, g.curNode))
	}
}

func (g *Gen) genRules(selector string, node ast.Selector) string {
	var css strings.Builder
	css.WriteString(selector + "{")

	for _, rule := range node.Rules {
		css.WriteString(rule.Name + ":" + rule.Value + ";")
	}

	for _, nested := range node.Nested {
		g.Output.WriteString(g.genRules(selector+" "+nested.SelectorText, nested))
	}

	css.WriteString("}")
	return css.String()
}
