package ast

import (
	"github.com/Krishanu230/Flipbook-Language/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

//// TODO: Add expression parsing support
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

//for debugging puposes
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

//implement the expression interface
func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

type StringLiteral struct {
	Token token.Token
	Value string
}

//implement the expression interface
func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }

type NewStatement struct {
	Token     token.Token
	Name      *Identifier
	DType     *Identifier
	DimX      *IntegerLiteral
	DimY      *IntegerLiteral
	Attribute *Identifier
}

//implement the statement interface
func (ns *NewStatement) statementNode()       {}
func (ns *NewStatement) TokenLiteral() string { return ns.Token.Literal }

type SetStatement struct {
	Token    token.Token
	Target   *Identifier
	Property *Identifier
	Value    *IntegerLiteral
}

func (ss *SetStatement) statementNode()       {}
func (ss *SetStatement) TokenLiteral() string { return ss.Token.Literal }

type InsertStatement struct {
	Token     token.Token
	Image     *Identifier
	Book      *Identifier
	StartPage *IntegerLiteral
	EndPage   *IntegerLiteral
	PositionX *IntegerLiteral
	PositionY *IntegerLiteral
}

func (is *InsertStatement) statementNode()       {}
func (is *InsertStatement) TokenLiteral() string { return is.Token.Literal }

type KeyframeStatement struct {
	Token         token.Token
	Image         *Identifier
	Book          *Identifier
	Property      *Identifier
	StartPage     *IntegerLiteral
	StartProperty *IntegerLiteral
	EndPage       *IntegerLiteral
	EndProperty   *IntegerLiteral
}

func (kfs *KeyframeStatement) statementNode()       {}
func (kfs *KeyframeStatement) TokenLiteral() string { return kfs.Token.Literal }

type SaveStatement struct {
	Token      token.Token
	Book       *Identifier
	OutputName *Identifier
}

func (svs *SaveStatement) statementNode()       {}
func (svs *SaveStatement) TokenLiteral() string { return svs.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
