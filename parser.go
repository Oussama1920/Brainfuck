package brainfuck

import (
	"io"
)

// RuneParser is to parse the tokens into instructions
type RuneParser interface {
	Parse() []*Instruction
}

// Instruction is the operation of our turing machine
// id the instruction
// additionalData is the position of the bracket of loops or number of occuerence for other operation
type Instruction struct {
	id             Identifier
	additionalData int
}

// Parser builds an abstract structure tree.
// Parser uses Stack to keep track of loops using push and pop methods
// s is the scanner to tokenerize the input
// buf is an internal struct to process input at a time of scan
// DesactivatedOperation will contain not supported operation for this instance

type Parser struct {
	s                     *Scanner
	instruction           []*Instruction
	buf                   Buf
	stack                 Stack
	DesactivatedOperation []DesactivatedOperation
}
type Buf struct {
	token  Identifier // last read token
	isUsed bool       // whether the token buffer is in use.
}
type DesactivatedOperation struct {
	operator string
}

// NewParser creates new Parser from input r.
func NewParser(r io.Reader) *Parser {
	desactivatedOperation := []DesactivatedOperation{}
	desactivatedOperation = append(desactivatedOperation, DesactivatedOperation{"²"})
	return &Parser{s: NewScanner(r), DesactivatedOperation: desactivatedOperation}
}

func (parser *Parser) Parse() []*Instruction {
	for {
		tok := parser.scan()
		if tok.Token == IllegalToken {
			break
		}
		switch tok.Token {
		case op_inc_dp:
			if parser.CheckActivation(">") {
				parser.AddInstruction(tok)
			}
		case op_dec_dp:
			if parser.CheckActivation("<") {
				parser.AddInstruction(tok)
			}
		case op_inc_val:
			if parser.CheckActivation("+") {
				parser.AddInstruction(tok)
			}
		case op_dec_val:
			if parser.CheckActivation("-") {
				parser.AddInstruction(tok)
			}
		case op_out:
			if parser.CheckActivation(".") {
				parser.AddInstruction(tok)
			}
		case op_in:
			if parser.CheckActivation(",") {
				parser.AddInstruction(tok)
			}
		case op_sqr_val:
			if parser.CheckActivation("²") {
				parser.AddInstruction(tok)
			}
		case op_jmp_fwd:
			openLoop := parser.BuildInstruction(tok, 0)
			parser.stack.Push(openLoop)
		case op_jmp_bck:
			openLoop := parser.stack.Pop().(int)
			closeLoop := parser.BuildInstruction(tok, openLoop)
			parser.instruction[openLoop].additionalData = closeLoop
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
		if next.Token != t.Token {
			p.unscan()
			break
		}
		c++
	}
	return p.BuildInstruction(t, c)
}

// buildInst creates a instruction from the given literals.
func (p *Parser) BuildInstruction(id Identifier, additionalData int) int {
	// build instruction
	instruction := &Instruction{
		id:             id,
		additionalData: additionalData,
	}
	// add inst to instruction list
	p.instruction = append(p.instruction, instruction)
	return len(p.instruction) - 1
}

// checkActivation checks if operator is active or no
// returns true if operator is activated
func (p *Parser) CheckActivation(operator string) bool {
	for _, identifier := range p.DesactivatedOperation {
		if identifier.operator == operator {
			return false
		}
	}
	return true
}

// Desactivate method is to desactivate operator
func (p *Parser) Desactivate(operator string) {
	p.DesactivatedOperation = append(p.DesactivatedOperation, DesactivatedOperation{operator})
}

// Activate method is to desactivate operator
func (p *Parser) Activate(operator string) {
	// will contain the index of the operator to delete
	index := 0
	// only for safe append
	counter := 0
	// check if already desactivated
	for i, identifier := range p.DesactivatedOperation {
		if identifier.operator == operator {
			index = i
			counter++
		}
	}
	if counter > 0 {
		p.DesactivatedOperation = append(p.DesactivatedOperation[:index], p.DesactivatedOperation[index+1:]...)

	}
}
