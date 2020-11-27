package parser

import (
	"Flipbook/ast"
	"Flipbook/lexer"
	"testing"
)

func TestNewStatements(t *testing.T) {
	input := `new book bookone = (10,20) 30;
  new image img = (20,30) krishanu.jpeg;
  `
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("Parsed programmed is nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("Parsed programmed len is not 2 it is %d", len(program.Statements))
	}
	test := []struct {
		expectedIdentifier string
	}{
		{"bookone"},
		{"img"},
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

func TestSetStatements(t *testing.T) {
	input := `set img scale 100;
  set img position (20,30);
  `
	l := lexer.NewLexer(input)
	p := NewParser(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("Parsed programmed is nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("Parsed programmed len is not 2 it is %d", len(program.Statements))
	}
	test := []struct {
		expectedIdentifier string
	}{
		{"scale"},
		{"position"},
	}
	for i, tt := range test {
		st := program.Statements[i]
		if !testSetStatement(t, st, tt.expectedIdentifier) {
			return
		}
	}
}

func testSetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "set" {
		t.Errorf("s.TokenLiteral is not 'new' it is %q", s.TokenLiteral())
		return false
	}
	newSt, ok := s.(*ast.SetStatement)
	if !ok {
		t.Errorf("s is not *ast.SetStatement it is %T", s)
		return false
	}
	if newSt.Property.TokenLiteral() != name {
		t.Errorf("newSt.Property is not '%s' it is %s", name, newSt.Property.TokenLiteral())
		return false
	}
	return true
}

//Helper Functions
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
