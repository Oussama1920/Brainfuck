package brainfuck

import (
	"io"
)

// RuneParser will parse tokens and pack them in instructions
// initial state of the RuneParser is Parse method
type RuneParser interface {
	Parse() []*Instruction
}

// Instruction is an abstraction for an operation which machine can understand
// id is one single instruction
// c is complementary information about instruction like position or counts of occurrence
type Instruction struct {
	id Identifier
	c  int
}

// Parser builds AST (abstract structure tree).
// Parser uses Stack to keep track of loops
// it contains the Scanner to tokenize the data from input
// buf is an internal struct to process input at a time of scan
// inst is an slice, which every member is one single instruction
type Parser struct {
	s           *Scanner
	instruction []*Instruction
	buf         Buf
	stack       Stack
}
type Buf struct {
	token  Identifier // last read token
	isUsed bool       // whether the token buffer is in use.
}

// NewParser creates new Parser from input r.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (parser *Parser) Parse() []*Instruction {
	for {
		tok := parser.scan()
		if tok.token == IllegalToken {
			break
		}
		switch tok.token {
		case
			op_inc_dp,
			op_dec_dp,
			op_inc_val,
			op_dec_val,
			op_out,
			op_in:
			parser.AddInstruction(tok)
		case op_jmp_fwd:
			openLoop := parser.BuildInstruction(tok, 0)
			parser.stack.Push(openLoop)
		case op_jmp_bck:
			openLoop := parser.stack.Pop().(int)
			closeLoop := parser.BuildInstruction(tok, openLoop)
			parser.instruction[openLoop].c = closeLoop
		}

	}
	return parser.instruction
}

// scan returns next token unit.
func (p *Parser) scan() Identifier {
	// there is a token on the buffer
	if p.buf.isUsed {
		p.buf.isUsed = false
		return p.buf.token
	}
	// read the next token from s
	token := p.s.Scan()
	p.buf.token = token
	return token
}

// unscan sends the already consumed token back to buff.
func (p *Parser) unscan() {
	p.buf.isUsed = true
}

// addInst adds instructions to []*inst of Parser
// for efficiency, if there are multiple occurrences of the
// same token consecutively, we will fold it.
func (p *Parser) AddInstruction(t Identifier) int {
	// token occurrence count
	c := 1
	for {
		next := p.scan()
		if next.token != t.token {
			p.unscan()
			break
		}
		c++
	}
	return p.BuildInstruction(t, c)
}

// buildInst creates a instruction from the given literals.
func (p *Parser) BuildInstruction(id Identifier, c int) int {
	// build instruction
	instruction := &Instruction{
		id: id,
		c:  c,
	}
	// add inst to instruction list
	p.instruction = append(p.instruction, instruction)
	return len(p.instruction) - 1
}
