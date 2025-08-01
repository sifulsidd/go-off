package ast

import "go-work/token"

// every node in our AST has to implement Node interface
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	// basically inheriting Node, so Statement also has TokenLiteral() function in it
	Node

	statementNode()
}

type Expression interface {
	Node

	expressionNode()
}

type Program struct {
	// every monkey program is a series of statements
	Statements []Statement
}

// essentially reads the first statement, if it is "let" it let's us know what type of statement this is 
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token

	// holds identities of binding value
	Name *Identifier

	// value for expression that produces the value
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token

	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}