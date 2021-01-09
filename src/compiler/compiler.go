package compiler

import (
	"heron/src/ast"
	"heron/src/lexer"
	"heron/src/parser"
	"os"
	"strings"
)

//Compiler represents the compiler
type Compiler struct {
	Program *ast.Program
	Output  strings.Builder
	curNode ast.Selector
}

//Compile generates the corresponding CSS from the AST
func (g *Compiler) Compile() {
	for _, selector := range g.Program.Rules {
		g.curNode = selector

		g.Output.WriteString(g.compileRules(g.curNode.SelectorText, g.curNode))
	}
}

func (g *Compiler) compileRules(selector string, node ast.Selector) string {
	var css strings.Builder
	css.WriteString(selector + "{")

	for _, rule := range node.Rules {
		css.WriteString(rule.Name + ":" + rule.Value + ";")
	}

	for _, nested := range node.Nested {
		g.Output.WriteString(g.compileRules(selector+" "+nested.SelectorText, nested))
	}

	css.WriteString("}")
	return css.String()
}

//Compile is a helper function that will use all packages to compile Heron code
func Compile(input []byte) strings.Builder {
	p := parser.New(lexer.New(input), os.Args[1])
	tree := p.ParseProgram()

	generator := &Compiler{Program: tree}
	generator.Compile()

	return generator.Output
}
