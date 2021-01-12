package parser

import (
	"github.com/poseidoncoder/heron/src/ast"
	"github.com/poseidoncoder/heron/src/lexer"
	"io/ioutil"
	"testing"
)

func TestImports(t *testing.T) {
	input := `
		@import 'test_import.he';
	`

	createParserTest(t, input, []ast.Selector{})
}

func TestSpaces(t *testing.T) {
	input := `
	ul {
		transition: color 1s;
	}

	ul:hover {
		color: red;
	}`

	expected := []ast.Selector{
		{
			SelectorText: "ul",
			Rules: []ast.Rule{
				{
					Name:  "transition",
					Value: " color 1s",
				},
			},
			Nested: []ast.Selector{},
		},
		{
			SelectorText: "ul:hover",
			Rules: []ast.Rule{
				{
					Name:  "color",
					Value: " red",
				},
			},
		},
	}

	createParserTest(t, input, expected)
}

func TestImport(t *testing.T) {
	createImportTest(t, `
		@import "test_import.he";

		p {
			color: blue;
		}
	`, []ast.Selector{
		{
			SelectorText: "p",
			Rules: []ast.Rule{
				{
					Name: "color",
					Value: " blue",
				},
			},
		},
	}, []ast.Selector{
		{
			SelectorText: "button",
			Rules: []ast.Rule{
				{
					Name:  "background-color",
					Value: " blue",
				},
			},
		},
	}, "test_import.he")
}

func createImportTest(t *testing.T, input string, expected []ast.Selector, expectedImport []ast.Selector, importName string) {
	p := createParserTest(t, input, expected)

	if len(p.Imports) == 0 {
		t.Error("not enough imported files")
	}

	file, err := ioutil.ReadFile(importName)
	if err != nil {
		t.Error("there was an error opening that imported file")
	}

	createParserTest(t, string(file), expectedImport)
}

func createParserTest(t *testing.T, input string, expected []ast.Selector) *ast.Program {
	l := lexer.New([]byte(input))
	p := New(l, "test")

	tree := p.ParseProgram()

	for i, node := range tree.Rules {
		testSelector(t, tree, expected, i, node)
	}

	return p.program
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
