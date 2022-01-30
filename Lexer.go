package brainfuck

import (
	"bufio"
	"io"
)

// LexReader is an interface that contain our  read() method.
type LexicalReader interface {
	read() rune
}

// Scanner impliment a token reader.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner gives  an instance of the Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		bufio.NewReader(r),
	}
}

// Read method return the next rune from the reader and return EOF caracter if no more runes.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return EOF
	}
	return ch
}

// Scan method returns the next Token from the reader.
func (s *Scanner) Scan() Identifier {

	// read next rune
	ch := s.read()

	return s.next(ch)
}

// next method is used in scanning
func (s *Scanner) next(ch rune) Identifier {
	// Check against individual code points next.
	switch ch {
	case '>':
		return Identifier{token: op_inc_dp, Value: string(ch)}
	case '<':
		return Identifier{token: op_dec_dp, Value: string(ch)}
	case '+':
		return Identifier{token: op_inc_val, Value: string(ch)}
	case '-':
		return Identifier{token: op_dec_val, Value: string(ch)}
	case '[':
		return Identifier{token: op_jmp_fwd, Value: string(ch)}
	case ']':
		return Identifier{token: op_jmp_bck, Value: string(ch)}
	case '.':
		return Identifier{token: op_out, Value: string(ch)}
	case ',':
		return Identifier{token: op_in, Value: string(ch)}
	default:
		return Identifier{token: IllegalToken, Value: "<nil>"}
	}
}
