package ast

import (
	"bytes"
	"monkey/token"
	"strings"
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

//BlockStatement => { <statements> }
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

// TokenLiteral is the token string
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

//StringLiteral => "hello world"; eg.
type StringLiteral struct {
	Token token.Token
	Value string
}

// TokenLiteral is the integer as a string
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// FunctionLiteral represents a function
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

//TokenLiteral is the token string
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

//ArrayLiteral is an array, mixed types are ok
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

//TokenLiteral is the token string
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

//PrefixExpression => !<expression> or -<expression>
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

//TokenLiteral is the prefix as a string
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression => <expression> + <expression> eg.
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

// TokenLiteral is the infix operator as a string
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean => true or false
type Boolean struct {
	Token token.Token
	Value bool
}

// TokenLiteral is string value of the token
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) expressionNode()      {}
func (b *Boolean) String() string       { return b.Token.Literal }

// Null => null
type Null struct {
	Token token.Token
}

// TokenLiteral is string value of the token
func (n *Null) TokenLiteral() string { return n.Token.Literal }
func (n *Null) expressionNode()      {}
func (n *Null) String() string       { return n.Token.Literal }

//IfExpression => if <condition> {<consequence} else {<Alternative>}
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

// TokenLiteral is string value of the token
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

//WhileExpression is a loop
type WhileExpression struct {
	Token token.Token
	Test  Expression
	Body  *BlockStatement
}

// TokenLiteral is string value of the token
func (we *WhileExpression) TokenLiteral() string { return we.Token.Literal }
func (we *WhileExpression) expressionNode()      {}
func (we *WhileExpression) String() string {
	var out bytes.Buffer

	out.WriteString("while")
	out.WriteString("( ")
	out.WriteString(we.Test.String())
	out.WriteString(" ) {\n")
	out.WriteString(we.Body.String())
	out.WriteString("\n}")

	return out.String()
}

//CallExpression is the brackets after a function
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

// TokenLiteral is string value of the token
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

// IndexExpression is for indexing arrays: <expression>[<expression>]
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

// TokenLiteral is string value of the token
func (ie *IndexExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IndexExpression) expressionNode()      {}
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}

//HashLiteral is the dictonary
type HashLiteral struct {
	Token token.Token
	Pairs map[Expression]Expression
}

// HashLiteral is string value of the token
func (hl *HashLiteral) TokenLiteral() string { return hl.Token.Literal }
func (hl *HashLiteral) expressionNode()      {}
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
