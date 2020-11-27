package parser

import (
	"Flipbook/ast"
	"Flipbook/lexer"
	"testing"
)

func TestNewStatements(t *testing.T) {
	input := `new book bookone = (10,20) 30;
  `
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("Parsed programmed is nil")
	}
	if len(program.Statements) != 1 {
		t.Fatalf("Parsed programmed len is not 2 it is %d", len(program.Statements))
	}
	test := []struct {
		expectedIdentifier string
	}{
		{"bookone"},
	}
	for i, tt := range test {
		st := program.Statements[i]
		if !testNewStatement(t, st, tt.expectedIdentifier) {
			return
		}
	}
}

func testNewStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "new" {
		t.Errorf("s.TokenLiteral is not 'new' it is %q", s.TokenLiteral())
		return false
	}
	newSt, ok := s.(*ast.NewStatement)
	if !ok {
		t.Errorf("s is not *ast.NewStatement it is %T", s)
		return false
	}
	if newSt.Name.Value != name {
		t.Errorf("newSt.Name.Vale is not '%s' it is %s", name, newSt.Name.Value)
		return false
	}
	if newSt.Name.TokenLiteral() != name {
		t.Errorf("s.Name is not '%s' it is %s", name, newSt.Name)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	if len(p.errors) == 0 {
		return
	}
	t.Errorf("Parsing has %d Errors", len(p.errors))
	for _, msg := range p.errors {
		t.Errorf("Parser error: %q", msg)
	}
	t.FailNow()
}
