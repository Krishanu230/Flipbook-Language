package lexer

import (
	"Flipbook/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `new image img = krishanu.jpg;
  new book b = 30 30
  set b pageCount 10
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEW, "new"},
		{token.IMAGE, "image"},
		{token.IDN, "img"},
		{token.ASSIGN, "="},
		{token.IDN, "krishanu.jpg"},
		{token.SEMICOLON, ";"},
		{token.NEW, "new"},
		{token.BOOK, "book"},
		{token.IDN, "b"},
		{token.ASSIGN, "="},
		{token.INT, "30"},
		{token.INT, "30"},
		{token.SET, "set"},
		{token.IDN, "b"},
		{token.PAGECNT, "pageCount"},
		{token.INT, "10"},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] Wrong Token type expected=%q instrad of %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] Wrong Literal expected=%q instead of %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
