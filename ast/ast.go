package ast

// every node in our AST has to implement Node interface
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	// basically inheriting Node, so Statement also has TokenLiteral() function in it
	Node

	statmentNode()
}

type Expression interface {
	Node

	expressionNode()
}

type Program struct {
	// every monkey program is a series of statements
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}