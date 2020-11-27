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
	case token.SET:
		return p.parseSetStatement()
	case token.INSERT:
		return p.parseInsertStatement()
	case token.KEYFRAME:
		return p.parseKeyFrameStatement()
	case token.SAVE:
		return p.parseSaveStatement()
	default:
		return nil
	}
}

func (p *Parser) parseNewStatement() *ast.NewStatement {
	st := &ast.NewStatement{Token: p.curToken}

	if !(p.expectPeek(token.BOOK, token.IMAGE)) {
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

func (p *Parser) parseSetStatement() *ast.SetStatement {
	st := &ast.SetStatement{Token: p.curToken}

	if !(p.expectPeek(token.IDN)) {
		return nil
	}

	st.Object = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.SCALE, token.POSITION) {
		return nil
	}
	st.Property = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return st
}

func (p *Parser) parseInsertStatement() *ast.InsertStatement {
	st := &ast.InsertStatement{Token: p.curToken}

	if !(p.expectPeek(token.IDN)) {
		return nil
	}
	st.Image = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.IDN) {
		return nil
	}
	st.Book = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return st
}

func (p *Parser) parseKeyFrameStatement() *ast.KeyframeStatement {
	st := &ast.KeyframeStatement{Token: p.curToken}

	if !(p.expectPeek(token.IDN)) {
		return nil
	}
	st.Image = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.SCALE, token.POSITION) {
		return nil
	}
	st.Property = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return st
}

func (p *Parser) parseSaveStatement() *ast.SaveStatement {
	st := &ast.SaveStatement{Token: p.curToken}

	if !(p.expectPeek(token.BOOK)) {
		return nil
	}
	st.Book = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

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

func (p *Parser) expectPeek(tlist ...token.TokenType) bool {
	for _, allowedToken := range tlist {
		if p.peekTokenIs(allowedToken) {
			p.nextToken()
			return true
		}
	}
	p.peekError(tlist...)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(tlist ...token.TokenType) {
	allowedList := ""
	for _, t := range tlist {
		allowedList += string(t) + ", "
	}
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", allowedList, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
