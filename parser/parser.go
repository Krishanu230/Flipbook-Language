package parser

import (
	"fmt"
	"strconv"

	"github.com/Krishanu230/Flipbook-Language/ast"
	"github.com/Krishanu230/Flipbook-Language/lexer"
	"github.com/Krishanu230/Flipbook-Language/token"
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
	//read two fill up the current and peek tokens
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

//specific statement parsers
func (p *Parser) parseNewStatement() *ast.NewStatement {
	st := &ast.NewStatement{Token: p.curToken}

	if !(p.expectPeek(token.BOOK, token.IMAGE)) {
		return nil
	}
	st.DType = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.IDN) {
		return nil
	}
	st.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	if !p.expectPeek(token.LBRACKET) {
		return nil
	}
	if !p.expectPeek(token.LBRACKET) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	val := p.parseIntegerLiteral()
	if val != nil {
		st.DimX = val
	} else {
		return nil
	}

	if !p.expectPeek(token.COMMA) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	val = p.parseIntegerLiteral()
	if val != nil {
		st.DimY = val
	} else {
		return nil
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	if !p.expectPeek(token.COMMA) {
		return nil
	}
	//if the new object is book
	if st.DType.Token.Type == token.BOOK {
		if !p.expectPeek(token.INT) {
			return nil
		}
		st.Attribute = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else {
		//if the new object is image
		if !p.expectPeek(token.FILENAME) {
			return nil
		}
		st.Attribute = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return st
}

func (p *Parser) parseSetStatement() *ast.SetStatement {
	st := &ast.SetStatement{Token: p.curToken}

	if !(p.expectPeek(token.IDN)) {
		return nil
	}
	st.Target = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.SCALE, token.POSITIONX) {
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

	if !p.expectPeek(token.FROM) {
		return nil
	}
	if !p.expectPeek(token.PAGE) {
		return nil
	}

	if !p.expectPeek(token.INT) {
		return nil
	}
	val := p.parseIntegerLiteral()
	if val != nil {
		st.StartPage = val
	} else {
		return nil
	}

	if !p.expectPeek(token.TO) {
		return nil
	}

	if !p.expectPeek(token.INT) {
		return nil
	}
	val = p.parseIntegerLiteral()
	if val != nil {
		st.EndPage = val
	} else {
		return nil
	}

	if !p.expectPeek(token.AT) {
		return nil
	}

	if !p.expectPeek(token.LBRACKET) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	val = p.parseIntegerLiteral()
	if val != nil {
		st.PositionX = val
	} else {
		return nil
	}
	if !p.expectPeek(token.COMMA) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	val = p.parseIntegerLiteral()
	if val != nil {
		st.PositionY = val
	} else {
		return nil
	}
	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return st
}

func (p *Parser) parseKeyFrameStatement() *ast.KeyframeStatement {
	st := &ast.KeyframeStatement{Token: p.curToken}

	if !(p.expectPeek(token.IDN)) {
		return nil
	}
	st.Image = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !(p.expectPeek(token.IDN)) {
		return nil
	}
	st.Book = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.SCALE, token.POSITIONX, token.POSITIONY) {
		return nil
	}
	st.Property = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LBRACKET) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	val := p.parseIntegerLiteral()
	if val != nil {
		st.StartPage = val
	} else {
		return nil
	}
	if !p.expectPeek(token.COMMA) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}

	val = p.parseIntegerLiteral()
	if val != nil {
		st.StartProperty = val
	} else {
		return nil
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	if !p.expectPeek(token.TO) {
		return nil
	}

	if !p.expectPeek(token.LBRACKET) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	val = p.parseIntegerLiteral()
	if val != nil {
		st.EndPage = val
	} else {
		return nil
	}
	if !p.expectPeek(token.COMMA) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}

	val = p.parseIntegerLiteral()
	if val != nil {
		st.EndProperty = val
	} else {
		return nil
	}

	if !p.expectPeek(token.RBRACKET) {
		return nil
	}
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return st
}

func (p *Parser) parseSaveStatement() *ast.SaveStatement {
	st := &ast.SaveStatement{Token: p.curToken}

	if !(p.expectPeek(token.IDN)) {
		return nil
	}
	st.Book = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !(p.expectPeek(token.FILENAME)) {
		return nil
	}
	st.OutputName = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return st
}

func (p *Parser) parseIntegerLiteral() *ast.IntegerLiteral {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 0)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = int(value)
	return lit
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
