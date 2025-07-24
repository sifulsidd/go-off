package lexer

import (
	"testing"

	"go-work/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},

		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},

		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType{
			// %d is the i value, it is used for int values 
			// %q will put the expected type and token type in quotes to make it readable 
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}