package token

type TokenType string


type Token struct {

	Type TokenType

	Literal string

}

// define token types in the monkey language
const (
	ILLEGAL = "ILLEGAL" // signifies a character we don't know about 

	EOF = "EOF" // end of file, let's parser know we can stop

	// Identifiers + literals 
	IDENT = "IDENT" // add, foobar, x, y, ...

	INT = "INT"

	// Operators 
	ASSIGN = "="

	PLUS = "+"

	// Delimeters
	COMMA = ","

	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// Keywords 
	FUNCTION = "FUNCTION"

	LET = "LET"
)
