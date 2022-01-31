package brainfuck

import (
	"io"
)

// Memory capacity
const MemorySize int = 5000

// interface to write out the execution results
type IWriter interface {
	Write() io.Writer
}

// interface for an interpreter
// Run method executes created instructions by Parser
type Interpreter interface {
	IWriter
	Run() error
}

// BrainFuck is our implementation for the Interpreter
// parser is to build the instructions from the input
// writer is to contain the output
// reader is the input
// err is to contain errors in read or write
// Memory is to identify the size of the cell and position of it's cursor
type BrainFuck struct {
	parser *Parser
	writer io.Writer
	reader io.Reader
	buf    []byte
	ip     int
	err    error
	memory Memory
}
type Memory struct {
	cell   [MemorySize]int
	cursor int
}

// NewInterpreter creates new Interpreter instance .
func NewInterpreter(r io.Reader, w io.Writer, parser *Parser) *BrainFuck {
	return &BrainFuck{
		parser: parser,
		writer: w,
		reader: r,
		buf:    make([]byte, 1),
	}
}

// Run method executes the instructions
func (bf *BrainFuck) Run() error {
	instruction := bf.parser.Parse()
	for bf.ip < len(instruction) {
		switch instruction[bf.ip].id.Value {
		case "-":
			bf.decrement(instruction[bf.ip].additionalData)
		case "+":
			bf.increment(instruction[bf.ip].additionalData)
		case "<":
			bf.skate(-instruction[bf.ip].additionalData)
		case ">":
			bf.skate(instruction[bf.ip].additionalData)
		case ",":
			bf.read(instruction[bf.ip].additionalData)
		case ".":
			bf.write(instruction[bf.ip].additionalData)
		case "²":
			bf.sqr()
		case "[":
			if bf.val() == 0 {
				bf.goTo(instruction[bf.ip].additionalData)
				continue
			}
		case "]":
			if bf.val() != 0 {
				bf.goTo(instruction[bf.ip].additionalData)
				continue
			}
		}
		bf.ip++
	}

	return bf.err
}

// sqr method calculate the square of the value of the current cell in memorry
// value is modulo [255]
func (bf *BrainFuck) sqr() {
	bf.memory.cell[bf.cur()] = (bf.memory.cell[bf.cur()] * bf.memory.cell[bf.cur()]) % 255
}

// cur method returns the position of current cursor in the memory
func (bf *BrainFuck) cur() int {
	return bf.memory.cursor
}

// skate method moves the current cursor in the memory to given offset
func (bf *BrainFuck) skate(offset int) {
	bf.memory.cursor += offset
}

// goTo method forwards the cursor to position p.
func (bf *BrainFuck) goTo(p int) {
	bf.ip = p
}

// inc method increments the value of the current cell in memory by v.
// value is modulo [255]
func (bf *BrainFuck) increment(v int) {
	bf.memory.cell[bf.cur()] = (bf.memory.cell[bf.cur()] + v) % 255
}

// decrement method decrements the value of the current cell in memory by v
// value is modulo [255]
func (bf *BrainFuck) decrement(v int) {
	if bf.memory.cell[bf.cur()]-v >= 0 {
		bf.memory.cell[bf.cur()] -= v
	} else {
		bf.memory.cell[bf.cur()] = 256 + bf.memory.cell[bf.cur()] - v
	}
}

// val method returns current value of which cursor is pointing.
func (b *BrainFuck) val() int {
	return b.memory.cell[b.cur()]
}

// write method writes the value in current cell of the memory
func (bf *BrainFuck) write(times int) bool {
	bf.buf[0] = byte(bf.val())
	for i := 0; i < times; i++ {
		if _, err := bf.writer.Write(bf.buf); err != nil {
			bf.err = err
			return false
		}
	}
	return true
}

// read method reads input buf
func (bf *BrainFuck) read(times int) bool {
	for i := 0; i < times; i++ {
		if _, err := bf.reader.Read(bf.buf); err != nil {
			bf.err = err
			return false
		}
		bf.memory.cell[bf.cur()] = int(bf.buf[0])
	}
	return true
}
