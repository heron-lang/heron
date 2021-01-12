package parser

import (
	"fmt"
	"github.com/poseidoncoder/heron/src/ast"
	"github.com/poseidoncoder/heron/src/errors"
	"github.com/poseidoncoder/heron/src/lexer"
	"github.com/poseidoncoder/heron/src/token"
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
			err := p.parseAtRule()

			if err != nil {
				err.Print()
			}
		}
	}

	return p.program
}

func (p *Parser) parseAtRule() (err *errors.Error) {
	if err = p.expectPeekToken(token.IDENT); err != nil {
		return
	}

	switch p.curToken.Literal {
	case "import":
		if err = p.expectPeekToken(token.STRING); err != nil {
			return
		}

		absolutePath, pathError := filepath.Abs(path.Join(filepath.Dir(p.program.FileName), p.curToken.Literal))

		if pathError != nil {
			err = &errors.Error{
				Msg:  "we had some trouble parsing that path",
				Type: errors.ImportError,
				Loc:  p.curToken.Loc,
			}

			return
		}

		imported, fileError := ioutil.ReadFile(absolutePath)

		if fileError != nil {
			goError := fileError.Error() //check what Go has to say about the error
			err = &errors.Error{
				Msg:  fmt.Sprintf("we had some trouble fetching '%v'\n\t%v", p.curToken.Literal, goError[:len(goError)-1] /*removes trailing newline*/),
				Type: errors.ImportError,
				Loc:  p.curToken.Loc,
			}

			return
		}

		p2 := New(lexer.New(imported), absolutePath)
		p2.ParseProgram()

		p.program.Imports = append(p.program.Imports, *p2.program)

		if err = p.expectPeekToken(token.EOS); err != nil {
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
			err = &errors.Error{
				Type: errors.SyntaxError,
				Msg:  fmt.Sprintf("unexpected %v, expected selector text", p.curToken.Type),
				Loc:  p.curToken.Loc,
			}

			return
		}

		p.nextToken()
	}

	if err = p.expectCurToken(token.LBRACE); err != nil {
		return
	}

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
			err = &errors.Error{
				Type: errors.SyntaxError,
				Msg:  fmt.Sprintf("unexpected %v, expected a colon or closing brace", p.peekToken.Type),
				Loc:  p.peekToken.Loc,
			}

			return
		}

		if err != nil {
			return
		}
	}

	p.nextToken()

	return
}

func (p *Parser) parseRule() (rule ast.Rule, err *errors.Error) {
	rule.Name = p.curToken.Literal

	if err = p.expectPeekToken(token.COLON); err != nil {
		return
	}

	for p.peekToken.Type != token.EOS {
		if err = p.expectPeekToken(token.IDENT); err != nil {
			return
		}

		rule.Value += " " + p.curToken.Literal
	}

	p.nextToken() //SEMICOLON
	p.nextToken() //RULE NAME

	return
}

func (p *Parser) expectCurToken(expected token.TokenType) *errors.Error {
	return p.expectToken(p.curToken, expected)
}

func (p *Parser) expectPeekToken(expected token.TokenType) *errors.Error {
	return p.expectToken(p.peekToken, expected)
}

func (p *Parser) expectToken(tok token.Token, expected token.TokenType) *errors.Error {
	if tok.Type != expected {
		return &errors.Error{
			Msg: fmt.Sprintf("unexpected %v, expected %v", tok.Type, expected),
			Type: errors.SyntaxError,
			Loc: tok.Loc,
		}
	}

	p.nextToken()
	return nil
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
