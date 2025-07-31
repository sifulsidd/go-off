# Tree-walking Interpreter #

## Lexical Analysis ## 
Lexical analysis or lexing transoforms source code to tokens. 
Tokens are small, easily categorizable data structures that are fed to the parser.
Then the parser performs a second transofrmation and turns the token into an *"Abstract Syntax Tree"*

*Lexer* will take source code as input and output the tokens that represent the source code. 

When we see a *func (l *Lexer) methodName returnedValue* essentially what it's doing is creating a function in our class, python equivalent -> *def methodName(self) -> returnedValue:*, python keeps it in the same class wheras go has to make it outside of the struct class

### Lexer ####
Worked on the lexer, essentially it reads each character and determines what kind of value it contains. Whether it be an integer, keyword or even a symbol. We check it at each point and make sure it is reading what we are typing. Fully tested in TDD as well. 

#### REPL (read, eval, print, loop)/Console #### 
Starts off by >> and then reads your code
the first for loop without and conditions is basically the same as a python while True loop
*Start function:*
- Takes the input we provide
- loops until it is false
- reads the input the scanner is provided, if empty just exits
- if not empty we create a lexer and pass in the string we have
- goes through whole string until it reaches "EOF" and returns that value each time 


### Parser ### 
Parser - turns input into a data structure that represents input
Parsers take source code as an input and produce data structures which represents this source code

*Context-free grammar(CFG)* - set of rules that describe how to form correct (valid according to syntax) sentences in a language

Parser repeatedly advances tokens and checks current token to decide what to do next: either call another parsing function or throw an error.