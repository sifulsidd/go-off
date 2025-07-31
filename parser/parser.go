package parser

import (
	"go-work/token"
	"go-work/ast"
	"go-work/lexer"
)

// basically empty Parser class, think of lexer as a variable you pass in self.l = lexer
type Parser struct {
	// self.l = lexer
	l *lexer.Lexer
	// curToken = None
	curToken token.Token
	//peekToken = None
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	// calls parser, think of it as Parser(Lexer), but Lexer is represented by l 
	p := &Parser{l: l}
	
	// read two tokens, so curToken and peekToken are set, that's why we use nextToken twice one for cur and other for next
	p.nextToken()
	p.nextToken()
	return p
}


func (p *Parser) nextToken(){
	// current position becomes the the value that comes next, current constantly becomes peek and peek becomes the nextOne, it's a sliding window
	p.curToken = p.peekToken
	// we make the peek move by finding the lexer and going to the next token 
	p.peekToken = p.l.NextToken()
}

// just a placeholder will have functionality soon 
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}