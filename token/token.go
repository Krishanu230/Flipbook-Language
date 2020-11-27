package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDN    = "IDN"
	INT    = "INT"
	IMAGE  = "IMAGE"
	PAGE   = "PAGE"
	BOOK   = "BOOK"
	STRING = "STRING"

	//operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LBRACKET = "("
	RBRACKET = "("
	LBRACE   = "{"
	RBRACE   = "}"

	//KeyWords
	EFFECT  = "EFFECT"
	NEW     = "NEW"
	AT      = "AT"
	IN      = "IN"
	SET     = "SET"
	PAGECNT = "PAGECNT"
)

var keywords = map[string]TokenType{
	"effect":    EFFECT,
	"new":       NEW,
	"at":        AT,
	"in":        IN,
	"set":       SET,
	"pageCount": PAGECNT,
	"image":     IMAGE,
	"book":      BOOK,
}

func LookupIdent(iden string) TokenType {
	if tok, ok := keywords[iden]; ok {
		return tok
	}
	return IDN
}
