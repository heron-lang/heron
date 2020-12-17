package parser

import (
	"ares/src/ast"
	"ares/src/scanner"
	"ares/src/token"
	"errors"
	"fmt"
	"os"
)

type Parser struct {
	l *scanner.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *scanner.Lexer) *Parser {
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

func (p *Parser) parseSelector() (selector ast.Selector, error error) {
	if p.curToken.Type == token.IDENT {
		//SELECTOR TEXT
		selector.SelectorText = p.curToken.Literal

		p.nextToken() //{

		if p.curToken.Type == token.LBRACE {
			p.nextToken()                         //RULE NAME (skips LBRACE)
			for p.curToken.Type != token.RBRACE { //Parse rule until it meets }
				rule, err := p.parseRule()

				if err != nil {
					error = err
					return
				}

				selector.Rules = append(selector.Rules, rule)
			}
		} else {
			error = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected opening brace", p.curToken.Type))
			return
		}
	} else {
		error = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected selector", p.curToken.Type))
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
			p.nextToken() //SKIP RULE VALUE
			p.nextToken() //SKIP SEMI-COLON
		} else {
			err = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected rule value", p.curToken.Type))
		}
	} else {
		err = errors.New(fmt.Sprintf("Syntax Error: unexpected %v, expected rule name", p.curToken.Type))
	}

	return
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
