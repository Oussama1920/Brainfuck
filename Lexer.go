package brainfuck

import (
	"bufio"
	"io"
)

// LexReader is an interface that wraps Read method.
// Read method reads and return the next rune from the input.
type LexReader interface {
	read() rune
}

// LexScanner is the interface that adds Unread method to the
// basic LexReader.
//
// Unread causes the next call to the Read method return the same
// rune as the same previous call to Read.
type LexScanner interface {
	LexReader
	unread() error
}

// Scanner implements a tokenizer.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{
		bufio.NewReader(r),
	}
}

// Read method reads the next rune from r.
// err != nil only if there is no more rune to read.
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return EOF
	}
	return ch
}

// unread re-buffers the last read data.
func (s *Scanner) unread() error {
	if err := s.r.UnreadRune(); err != nil {
		return err
	}
	return nil
}

// Scan prepare and returns the next Token.
func (s *Scanner) Scan() Identifier {

	// read next rune
	ch := s.read()

	return s.next(ch)
}

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
