package brainfuck

import (
	"io"
)

// interface to write out the execution results
type Writer interface {
	Write() io.Writer
}

// interface for an interpreter
// Run method executes created instructions by Parser
type Interpreter interface {
	Writer
	Run() error
}

// Memory capacity
const MemorySize int = 3000

// BrainFuck is an implementation of the Interpreter
// it has internal parser which builds instructions from the input
// result is written into w
// memory struct keeps memory data and cursor to move between memory cells and update their values
// err != nil if any error happen during the print/read operation
type BrainFuck struct {
	p      *Parser
	w      io.Writer
	i      io.Reader
	buf    []byte
	ip     int
	err    error
	memory Memory
}
type Memory struct {
	cell   [MemorySize]int
	cursor int
}

// NewInterpreter creates new Interpreter instance and initialize it's internal Parser.
func NewInterpreter(i io.Reader, w io.Writer, parser *Parser) *BrainFuck {
	return &BrainFuck{
		p:   parser,
		w:   w,
		i:   i,
		buf: make([]byte, 1),
	}
}

// Run method executes the instructions
// err != nil if error happen during read/print operations
// output returns in format of bytes
func (b *BrainFuck) Run() error {
	inst := b.p.Parse()
	for b.ip < len(inst) {
		switch inst[b.ip].id.Value {
		case ">":
			b.seek(inst[b.ip].c)
		case "<":
			b.seek(-inst[b.ip].c)
		case "+":
			b.inc(inst[b.ip].c)
		case "-":
			b.dec(inst[b.ip].c)
		case ".":
			b.write(inst[b.ip].c)
		case ",":
			b.read(inst[b.ip].c)
		case "[":
			if b.val() == 0 {
				b.jump(inst[b.ip].c)
				continue
			}
		case "]":
			if b.val() != 0 {
				b.jump(inst[b.ip].c)
				continue
			}
		}
		b.ip++
	}

	return b.err
}

// curr method returns the position of current cursor in the memory
func (b *BrainFuck) cur() int {
	return b.memory.cursor
}

// seek method moves the cursor in the memory to given offset
// this move is relative to current cursor position
func (b *BrainFuck) seek(offset int) {
	b.memory.cursor += offset
}

// jump method forwards the cursor to position p.
func (b *BrainFuck) jump(p int) {
	b.ip = p
}

// reset method resets the cursor and writer to point to invalid state.
func (b *BrainFuck) reset() {
	b.memory.cursor = 0
}

// inc method increments the value of the current cell in memory by v.
func (b *BrainFuck) inc(v int) {
	b.memory.cell[b.cur()] = (b.memory.cell[b.cur()] + v) % 255
}

// dec method decrements the value of the current cell in memory by v.
func (b *BrainFuck) dec(v int) {
	if b.memory.cell[b.cur()]-v >= 0 {
		b.memory.cell[b.cur()] -= v
	} else {
		b.memory.cell[b.cur()] = 256 + b.memory.cell[b.cur()] - v
	}
}

// val method returns current value of which cursor is pointing.
func (b *BrainFuck) val() int {
	return b.memory.cell[b.cur()]
}

// doPrint method prints the value in current cell of the memory
// if any error happen during the Write operation err property will be set.
func (b *BrainFuck) write(times int) bool {
	b.buf[0] = byte(b.val())
	for i := 0; i < times; i++ {
		if _, err := b.w.Write(b.buf); err != nil {
			b.err = err
			return false
		}
	}
	return true
}

// doRead reads input from io
// if any error happen during the Read operation err property will be set.
func (b *BrainFuck) read(times int) bool {
	for i := 0; i < times; i++ {
		if _, err := b.i.Read(b.buf); err != nil {
			b.err = err
			return false
		}
		b.memory.cell[b.cur()] = int(b.buf[0])
	}
	return true
}
