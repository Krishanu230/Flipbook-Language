package lexer

import (
	"testing"

	"github.com/Krishanu230/Flipbook-Language/token"
)

func TestNextToken(t *testing.T) {
	input := `new book bookone (10,20) 30;
  new image img (20,30) krishanu.jpg;
  set img scale 10;
  `

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEW, "new"},
		{token.BOOK, "book"},
		{token.IDN, "bookone"},
		{token.LBRACKET, "("},
		{token.INT, "10"},
		{token.COMMA, ","},
		{token.INT, "20"},
		{token.RBRACKET, ")"},
		{token.INT, "30"},
		{token.SEMICOLON, ";"},
		{token.NEW, "new"},
		{token.IMAGE, "image"},
		{token.IDN, "img"},
		{token.LBRACKET, "("},
		{token.INT, "20"},
		{token.COMMA, ","},
		{token.INT, "30"},
		{token.RBRACKET, ")"},
		{token.IDN, "krishanu.jpg"},
		{token.SEMICOLON, ";"},
		{token.SET, "set"},
		{token.IDN, "img"},
		{token.SCALE, "scale"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.SWIRL, "swirl"},
		{token.IDN, "bookone"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
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
