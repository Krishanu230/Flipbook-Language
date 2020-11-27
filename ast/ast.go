package ast

import (
	"Flipbook/token"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type NewStatement struct {
	Token token.Token
	Name  *Identifier
	DType *Identifier
	Value Expression
}

func (ns *NewStatement) statementNode()       {}
func (ns *NewStatement) TokenLiteral() string { return ns.Token.Literal }

type SetStatement struct {
	Token    token.Token
	Object   *Identifier
	Property *Identifier
	Value    Expression
}

func (ss *SetStatement) statementNode()       {}
func (ss *SetStatement) TokenLiteral() string { return ss.Token.Literal }

type InsertStatement struct {
	Token     token.Token
	Image     *Identifier
	Book      *Identifier
	StartPage Expression
	EndPage   Expression
	Position  Expression
}

func (is *InsertStatement) statementNode()       {}
func (is *InsertStatement) TokenLiteral() string { return is.Token.Literal }

type KeyframeStatement struct {
	Token         token.Token
	Image         *Identifier
	Property      *Identifier
	StartPage     Expression
	StartProperty Expression
	EndPage       Expression
	EndProperty   Expression
}

func (kfs *KeyframeStatement) statementNode()       {}
func (kfs *KeyframeStatement) TokenLiteral() string { return kfs.Token.Literal }

type SaveStatement struct {
	Token token.Token
	Book  *Identifier
	//OutputPath *Identifier
}

func (svs *SaveStatement) statementNode()       {}
func (svs *SaveStatement) TokenLiteral() string { return svs.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
