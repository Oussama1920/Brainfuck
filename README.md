

Package `brainfuck` implements a brain fuck interpreter.

#### Basics
Interpreting occurs in three steps. 

First, the `Lexer` breaks up a stream of code
points (runes) into tokens.

These tokens represent the basic units of brain fuck syntax tree, such as whitespace, identifiers (like: > < + - . ,) and loops ( [ ] ).

Each scan returns a single token.

types: 
```go 
	IllegalToken       
	op_dec_dp             // <
	op_inc_dp            // >
	op_inc_val             // +
	op_dec_val            // -
	op_out            // .
	op_in             // ,
	op_jmp_fwd      // [
	RightBraop_jmp_bckcketToken     // ]
	WhitespaceToken      
	
```

- The second step is to feed these tokens into the parser which creates the abstract syntax tree (AST) based on
the context of the tokens. below you will see a simple AST.

```go
{
    {t:Token{Tok:op_dec_val, Value:"-"}, c:1},
    {t:Token{Tok:op_jmp_fwd, Value:"["}, c:7},
    {t:Token{Tok:op_dec_val, Value:"-"}, c:2},
    {t:Token{Tok:op_jmp_fwd, Value:"["}, c:5},
    {t:Token{Tok:PlusToken, Value:"+"}, c:1},
    {t:Token{Tok:op_jmp_bck, Value:"]"}, c:3},
    {t:Token{Tok:op_dec_val, Value:"-"}, c:2},
    {t:Token{Tok:op_jmp_bck, Value:"]"}, c:1},
}
```
- The last step is to execute the produced instructions in step two by interpreter.

Abstract syntax tree, is very simple but not flat. The first level is simply contains
identifiers (like: > < + - . ,). The second level contains the loops and it's internal
blocks ( which can contains identifiers and loop block again), etc.

#### How to use

	// create new io.Reader from inputs
	code := strings.NewReader("----[---->+<]>++.+.+.+.")
	
	// initialize the Parser with input
	parser := brainfuck.NewParser(code)
	
	// Standards interface to io
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)
	
	// initialize the machine
	bfm := brainfuck.NewInterpreter(input, output, parser)
	
    // Store the result in output interface 
	_ = bfm.Run()
	
	// print the result 
	fmt.Println (output.String())
	
#### How to run tests

To run all tests in the root of the project run `go test ./...` command.

#### 
