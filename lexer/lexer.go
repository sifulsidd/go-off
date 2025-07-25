package lexer 

import "go-work/token"

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)

	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case ',':
		tok = newToken(token.COMMA, l.ch)
	
	case '+':
		tok = newToken(token.PLUS, l.ch)
	
	case 0:
		tok.Literal = ""

		tok.Type = token.EOF
	}

	l.readChar()

	return tok

}

// look at current character under examination and return token depending on which character it is 
// essentially turns characters into tokens 
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// struct is same as a class in python 
/*
Keeps track of what the source code is,
Knows where it is in the input (position),
Knows where it's about to look (readPosition),
And holds the current character it's examining (ch).
*/
type Lexer struct {
	input string
	position int // current position in input (points to current character)
	readPosition int //current reading posiiton in input after current character
	ch byte // current character under examination
}

// create and run a new instance of Lexer
// go needs this because it doesn't have __init__ method from python
func New(input string) *Lexer{
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// gives us next character and advance our position in input string
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // sets l.ch to NUL, means end of input
	} else {
		l.ch = l.input[l.readPosition] // not end of input so continue to access next file
	}
	l.position = l.readPosition // position gets updated to the readPosition 
	l.readPosition += 1 // readPostion always points to next position
}



