package ast

import (
	"bytes"
	"monkey/token"
)

// Node is the base interface for all nodes in the parser
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String gets the literal string of the LetStatement
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement => return <expression>
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is the token string
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement => <expression>
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral is the token string of the first token in the expression
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Identifier is the Node for variable names
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral is the token string
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) expressionNode()      {}
func (i *Identifier) String() string       { return i.Value }

//IntegerLiteral => 5; eg.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral is the integer as a string
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) String() string       { return il.Token.Literal }
