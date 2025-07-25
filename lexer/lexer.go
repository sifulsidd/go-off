package lexer

import "go-work/token"

func (l *Lexer) NextToken() token.Token {
	// just an instance of a token value
	// we use token.Token because token helps us find the token package and then that struct
	// if we had token in this packacge we could've just wrote "var tok Token"
	var tok token.Token

	l.skipWhitespace()

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

	default:
		// reads the whole string until it gets to end of string and sets tok.Literal to the string
		// first checks character by character
		if isLetter(l.ch) {
			// if it is a character, reads the whole thing until it is empty "fn" , "let", etc.
			tok.Literal = l.readIdentifier()
			// if a token is referencing a valid token, we return that value or we return IDENT
			tok.Type = token.LookupIdent((tok.Literal))
			return tok
		} else if isDigit(l.ch) { //basically checking if token is an integer or not
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
			l.readChar()
			return tok
		}

	}

	l.readChar()

	return tok

}

// reads an identifier and keeps parsing until it reaches a non-letter character
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	// this is basically returning two indexes to give us the substring that we need for the string values
	return l.input[position:l.position]
}

// just checks whether given character is a string or not
func isLetter(ch byte) bool {
	// essentially this line allows our lexer to parse a string and accept the values below provided A-Z a-z ? ! _
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '!' || ch == '?'
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
	input        string
	position     int  // current position in input (points to current character)
	readPosition int  //current reading posiiton in input after current character
	ch           byte // current character under examination
}

// create and run a new instance of Lexer
// go needs this because it doesn't have __init__ method from python
func New(input string) *Lexer {
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
	l.readPosition += 1         // readPostion always points to next position
}

// basically skips any of the values we see below and if seen, moves onto the next character
func (l *Lexer) skipWhitespace() {
	// checking space, tabs, newlines or carriage returns
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// essentially gives me only the values that make up an integer
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	// position stores starting position and l.position is updated position
	return l.input[position:l.position]
}

// checks whether value is between 0 and 9
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
