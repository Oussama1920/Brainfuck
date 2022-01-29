package brainfuck

import (
	"bytes"
	"strings"
	"testing"
)


func TestBrainFuckMachine_LoopOperation(t *testing.T) {
	code := strings.NewReader("----[---->+<]>++.+.+.+.")
	parser := NewParser(code)
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)
	bfm := NewInterpreter(input, output, parser)
	_ = bfm.Run()

	if output.String() != "ABCD" {
		t.Errorf("wrong output, got %s", output.String())
	}

}

func TestBrainFuckMachine_PrintHelloWorld(t *testing.T) {
	input := `++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.`
	code := strings.NewReader(input)
	parser := NewParser(code)
	i := new(bytes.Buffer)
	o := new(bytes.Buffer)
	bfm := NewInterpreter(i, o, parser)
	_ = bfm.Run()

	if o.String() != "Hello World!\n" {
		t.Errorf("wrong output, got %s", o.String())
	}

}
