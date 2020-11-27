package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDN      = "IDN"
	INT      = "INT"
	FILENAME = "FILENAME"

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
	EFFECT   = "EFFECT"
	NEW      = "NEW"
	FROM     = "FROM"
	AT       = "AT"
	TO       = "TO"
	SET      = "SET"
	IMAGE    = "IMAGE"
	PAGE     = "PAGE"
	BOOK     = "BOOK"
	SCALE    = "SCALE"
	POSITION = "POSITION"
	INSERT   = "INSERT"
	KEYFRAME = "KEYFRAME"
	SAVE     = "SAVE"
)

var keywords = map[string]TokenType{
	"effect":   EFFECT,
	"new":      NEW,
	"at":       AT,
	"to":       TO,
	"set":      SET,
	"image":    IMAGE,
	"book":     BOOK,
	"scale":    SCALE,
	"insert":   INSERT,
	"from":     FROM,
	"page":     PAGE,
	"keyframe": KEYFRAME,
	"save":     SAVE,
	"position": POSITION,
}

func LookupIdent(iden string) TokenType {
	if tok, ok := keywords[iden]; ok {
		return tok
	}
	return IDN
}
