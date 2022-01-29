package brainfuck

const (
	IllegalToken Token = iota
	op_dec_dp          // <
	op_inc_dp          // >
	op_inc_val         // +
	op_dec_val         // -
	op_out             // .
	op_in              // ,
	op_jmp_fwd         // [
	op_jmp_bck         // ]

)

// special var to indicate end of stream.
var EOF = rune(-1)

// Tok represents a lexical token type.
type Token int

// Identifier represents a lexical tokens.
type Identifier struct {
	// the type of token.
	token Token

	// The literal value of the token(as parsed).
	Value string

	// The rune used for string tokens
	// Ending rune

	// Used for numeric tokens.
	// Number float64
	// Unit   string
}
