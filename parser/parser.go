package parser

import (
	"Flipbook/ast"
	"Flipbook/lexer"
	"Flipbook/token"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		st := p.parseStatement()
		if st != nil {
			program.Statements = append(program.Statements, st)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.NEW:
		return p.parseNewStatement()
	default:
		return nil
	}
}

func (p *Parser) parseNewStatement() *ast.NewStatement {
	st := &ast.NewStatement{Token: p.curToken}

	if !(p.expectPeek(token.BOOK)) {
		return nil
	}
	st.DType = &ast.Identifier{Token: p.curToken, Value: ""}
	if !p.expectPeek(token.IDN) {
		return nil
	}
	st.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return st
}

//Helper Functions
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
