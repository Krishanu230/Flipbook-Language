package lexer

import (
	"strings"

	"github.com/Krishanu230/Flipbook-Language/token"
)

type Lexer struct {
	input        string
	prevPosition int
	position     int
	ch           byte //char at prevPosition
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipSpace()
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ')':
		tok = newToken(token.RBRACKET, l.ch)
	case '(':
		tok = newToken(token.LBRACKET, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '"':
		fname := l.readFilename()
		if fname == "" {
			tok := newToken(token.ILLEGAL, l.ch)
			return tok
		}
		tok.Type = token.FILENAME
		tok.Literal = fname

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNum()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.ch = 0 //ASCI NUL
	} else {
		l.ch = l.input[l.position]
	}
	l.prevPosition = l.position
	l.position += 1
}

func (l *Lexer) readFilename() string {
	prevPos := l.prevPosition
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	fname := l.input[prevPos:l.prevPosition]
	split := strings.Split(fname, ".")
	if len(split) != 2 {
		return ""
	}
	if !(split[1] == "jpeg" || split[1] == "png" || split[1] == "jpg" || split[1] == "pdf") {
		return ""
	}
	return l.input[prevPos+1 : l.prevPosition]
}

func (l *Lexer) readIdentifier() string {
	prevPos := l.prevPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[prevPos:l.prevPosition]
}

func (l *Lexer) readNum() string {
	prevPos := l.prevPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[prevPos:l.prevPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func (l *Lexer) skipSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
