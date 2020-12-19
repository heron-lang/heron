package parser

import (
	"ares/src/ast"
	"ares/src/lexer"
	"testing"
)

func Test_ParseProgram(t *testing.T) {
	l := lexer.New(`
	ul {
		li {
			color: red;
		}

		background-color: blue;
	}`)

	p := New(l)
	tree := p.ParseProgram()

	expected := []ast.Selector{
		{
			SelectorText: "ul",
			Rules: []ast.Rule{
				{
					Name:  "background-color",
					Value: "blue",
				},
			},
			Nested: []ast.Selector{
				{
					SelectorText: "li",
					Rules: []ast.Rule{
						{
							Name:  "color",
							Value: "red",
						},
					},
				},
			},
		},
	}

	for i, node := range tree.Rules {
		testSelector(t, tree, expected, i, node)
	}
}

func testSelector(t *testing.T, tree *ast.Program, expected []ast.Selector, i int, node ast.Selector) {
	expectedNode := expected[i]
	if node.SelectorText != expectedNode.SelectorText {
		t.Errorf("Expected %v for the selector text but got %v instead", expectedNode.SelectorText, node.SelectorText)
	}

	for r, rule := range node.Rules {
		expectedRule := expectedNode.Rules[r]

		if expectedRule.Value != rule.Value {
			t.Errorf("Expected %v for the rule value but got %v instead", expectedRule.Value, rule.Value)
		}

		if expectedRule.Name != rule.Name {
			t.Errorf("Expected %v for the rule value but got %v instead", expectedRule.Name, rule.Name)
		}

		for n, nested := range node.Nested {
			testSelector(t, tree, expected[n].Nested, n, nested)
		}
	}
}
