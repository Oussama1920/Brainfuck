/*

Package brainfuck implements a brain fuck interpreter.

Basics
Interpreting occurs in three steps. First, the Lexer breaks up a stream of code
points (runes) into tokens. These tokens represent the units of brain fuck syntax tree
such as whitespace, identifiers (like: > < + - . ,) and loops ( [ ] ). The second step is to feed
these tokens into the parser which creates the abstract syntax tree (AST) based on
the context of the tokens.
The last step is to execute the produced instructions in step two by interpreter.

Abstract syntax tree, is very simple but not flat. The first level is simply contains
identifiers (like: > < + - . ,). The second level contains the loops and it's internal
blocks ( which can contains identifiers and loop block again).

*/
package brainfuck
