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

	MINUS = "-"

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
	// Boolean keywords 
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

	// Operators
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"
)

var keywords = map[string]TokenType{
	"fn" : FUNCTION,
	"let": LET,
	"true": TRUE, 
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

// checks if given identifier is actually a keyword or not, if it isn't return IDENT
func LookupIdent(ident string) TokenType {
	// essentially the ok value will have true or false, if it is true we return the tok, otherwise we return IDENT
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}

