package parser

import (
	"ares/src/ast"
	"ares/src/scanner"
	"testing"
)

func Test_ParseProgram(t *testing.T) {
	l := scanner.New(`
	p {
       color: red;
	}`)

	p := New(l)
	tree := p.ParseProgram()

	expected := []ast.Selector{
		{
			SelectorText: "p",
			Rules: []ast.Rule{
				{
					Name:  "color",
					Value: "red",
				},
			},
		},
	}

	for i, node := range tree.Rules {
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
		}
	}
}
