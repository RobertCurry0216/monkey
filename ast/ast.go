package ast

import (
	"monkey/token"
)

// Node is the base interface for all nodes in the parser
type Node interface {
	TokenLiteral() string
}

// Statement is not strictly neccessary but will be usefull to distinguish node types
type Statement interface {
	Node
	statementNode()
}

// Expression is not strictly necessary but will be usefull to distinguish node types
type Expression interface {
	Node
	expressionNode()
}

// Program will be the root node for all ast's
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal string of the node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is the Node for statements like: let x = 5;
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral is the token string
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier is the Node for variable names
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral is the token string
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
