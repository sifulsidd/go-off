# Tree-walking Interpreter #

## Lexical Analysis ## 
Lexical analysis or lexing transoforms source code to tokens. 
Tokens are small, easily categorizable data structures that are fed to the parser.
Then the parser performs a second transofrmation and turns the token into an *"Abstract Syntax Tree"*

*Lexer* will take source code as input and output the tokens that represent the source code. 

When we see a *func (l *Lexer) methodName returnedValue* essentially what it's doing is creating a function in our class, python equivalent -> *def methodName(self) -> returnedValue:*, python keeps it in the same class wheras go has to make it outside of the struct class