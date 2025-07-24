package lexer 

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

func New(input string) *Lexer{
	l := &Lexer{input: input}
	return l
}