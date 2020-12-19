package parser

import (
	"ares/src/ast"
	"ares/src/lexer"
	"ares/src/token"
	"errors"
	"fmt"
	"os"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	//read two so that curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Rules = []ast.Selector{} //makes sure that the rules are not null

	for p.curToken.Type != token.EOF {
		selector, err := p.parseSelector()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		program.Rules = append(program.Rules, selector)
		p.nextToken()
	}

	return program
}

func (p *Parser) parseSelector() (selector ast.Selector, err error) {
	for p.curToken.Type != token.LBRACE {
		if p.curToken.Type == token.COLON || p.curToken.Type == token.IDENT {
			selector.SelectorText += p.curToken.Literal
		} else {
			err = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected selector", p.curToken.Type))
			return
		}

		p.nextToken()
	}

	if p.curToken.Type == token.LBRACE {
		p.nextToken()                         //RULE NAME
		for p.curToken.Type != token.RBRACE { //Parse rule until it meets }
			//This is disturbing nested curly braces by cutting off too early
			switch p.peekToken.Type {
			case token.COLON:
				var rule ast.Rule
				rule, err = p.parseRule()
				selector.Rules = append(selector.Rules, rule)
			case token.LBRACE:
				var nested ast.Selector
				nested, err = p.parseSelector()
				selector.Nested = append(selector.Nested, nested)
			}

			if err != nil {
				return
			}
		}

		p.nextToken() //RULE NAME (skips RBRACE)
	} else {
		err = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected opening brace", p.curToken.Type))
		return
	}

	return
}

func (p *Parser) parseRule() (rule ast.Rule, err error) {
	if p.curToken.Type == token.IDENT {
		rule.Name = p.curToken.Literal

		p.nextToken() //COLON
		p.nextToken() //RULE VALUE

		if p.curToken.Type == token.IDENT {
			rule.Value = p.curToken.Literal
			p.nextToken() //SEMICOLON
			p.nextToken() //RULE NAME
		} else {
			err = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected rule value", p.curToken.Type))
		}
	} else {
		err = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected rule name", p.curToken.Type))
	}

	return
}

func (p *Parser) nextToken() {
	//fmt.Println(p.curToken)

	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
