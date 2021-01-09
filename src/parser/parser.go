package parser

import (
	"fmt"
	"heron/src/ast"
	"heron/src/errors"
	"heron/src/lexer"
	"heron/src/token"
	"io/ioutil"
	"path"
	"path/filepath"
)

//Parser represents the program parser
type Parser struct {
	l       *lexer.Lexer
	program *ast.Program

	curToken  token.Token
	peekToken token.Token
}

//New creates a new parser
func New(l *lexer.Lexer, fileName string) *Parser {
	p := &Parser{l: l, program: &ast.Program{FileName: fileName}}

	//read two so that curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

//ParseProgram creates the programs AST (Abstract Syntax Tree)
func (p *Parser) ParseProgram() *ast.Program {
	for p.curToken.Type != token.EOF {
		switch p.curToken.Type {
		case token.IDENT:
			selector, err := p.parseSelector()

			if err != nil {
				err.Print()
			}

			p.program.Rules = append(p.program.Rules, selector)
		case token.ATRULE:
			p.parseAtRule()
		}
	}

	return p.program
}

func (p *Parser) parseAtRule() (err *errors.Error) {
	p.nextToken() //IDENT

	if err = p.expectToken(token.IDENT); err != nil {
		return
	}

	switch p.curToken.Literal {
	case "import":
		p.nextToken() //STRING

		if err = p.expectToken(token.STRING); err != nil {
			return
		}

		absolutePath, pathError := filepath.Abs(path.Join(path.Dir(p.program.FileName), p.curToken.Literal))

		if pathError != nil {
			err = &errors.Error{
				Msg:  "we had some trouble parsing that path",
				Type: errors.ImportError,
				Loc:  p.curToken.Loc,
			}

			err.Print()
			return
		}

		imported, fileError := ioutil.ReadFile(absolutePath)

		if fileError != nil {
			err = &errors.Error{
				Msg:  fmt.Sprintf("we had some trouble fetching '%v'\n\t%v", p.curToken.Literal, fileError.Error()),
				Type: errors.ImportError,
				Loc:  p.curToken.Loc,
			}
			return
		}

		p2 := New(lexer.New(imported), p.curToken.Literal)
		p2.ParseProgram()

		p.program.Imports = append(p.program.Imports, *p2.program)

		p.nextToken() //EOS

		if err = p.expectToken(token.EOS); err != nil {
			return
		}

		p.nextToken() //SKIP OVER EOS
		break
	}

	return
}

func (p *Parser) parseSelector() (selector ast.Selector, err *errors.Error) {
	for p.curToken.Type != token.LBRACE {
		if p.curToken.Type == token.COLON || p.curToken.Type == token.IDENT {
			selector.SelectorText += p.curToken.Literal
		} else {
			err = p.newError(fmt.Sprintf("unexpected %v, expected selector", p.curToken.Type))
			return
		}

		p.nextToken()
	}

	if err = p.expectToken(token.LBRACE); err != nil {
		return
	}

	p.nextToken()                         //RULE NAME
	for p.curToken.Type != token.RBRACE { //Parse rule until it meets }
		switch p.peekToken.Type {
		case token.COLON:
			var rule ast.Rule
			rule, err = p.parseRule()
			selector.Rules = append(selector.Rules, rule)
		case token.LBRACE:
			var nested ast.Selector
			nested, err = p.parseSelector()
			selector.Nested = append(selector.Nested, nested)
		default:
			err = p.newError(fmt.Sprintf("unexpected %v, expected a colon or right brace", p.peekToken.Type))
			return
		}

		if err != nil {
			return
		}
	}

	p.nextToken() //RULE NAME (skips RBRACE)

	return
}

func (p *Parser) parseRule() (rule ast.Rule, err *errors.Error) {
	if err = p.expectToken(token.IDENT); err != nil {
		return
	}

	rule.Name = p.curToken.Literal

	p.nextToken() //COLON

	for p.peekToken.Type != token.EOS {
		p.nextToken() //RULE VALUE

		if err = p.expectToken(token.IDENT); err != nil {
			return
		}

		rule.Value += " " + p.curToken.Literal
	}

	p.nextToken() //SEMICOLON
	p.nextToken() //RULE NAME

	return
}

func (p *Parser) expectToken(expected token.TokenType) *errors.Error {
	if p.curToken.Type != expected {
		return p.newError(fmt.Sprintf("unexpected %v, expected %v", p.curToken.Type, expected))
	}

	return nil
}

func (p *Parser) newError(msg string) *errors.Error {
	return &errors.Error{
		Msg:  msg,
		Loc:  p.curToken.Loc,
		Type: errors.SyntaxError,
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
