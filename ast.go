package brainfuck

// the successive integer constants for each operation in golang
// IllegalToken is to identify any token out of brainFuck operations.
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

// EOF is for end of file (stream).
var EOF = rune(-1)

// Token represents a lexical for the type of token.
type Token int

// Identifier represents a lexical tokens.
type Identifier struct {
	// the type of the token.
	token Token

	// The parsed value of the token.
	Value string
}
