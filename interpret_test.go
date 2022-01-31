package brainfuck

import (
	"bytes"
	"strings"
	"testing"
)

func TestBrainFuckMachine(t *testing.T) {
	for _, tt := range []struct {
		name   string
		wanted string
		input  string
	}{
		{"Print Hello World", "Hello, World", HelloWordInBrainFuck},
		{"Print fibonacci", "1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89", FibonacciInBrainfFuck},
		{"Print factorial", "0! = 1\n1! = 1\n2! = 2\n3! = 6\n4! = 24\n5! = 120\n6! = 210\n", FactorialInBrainFuck},
	} {
		code := strings.NewReader(tt.input)
		parser := NewParser(code)
		input := new(bytes.Buffer)
		output := new(bytes.Buffer)
		bfm := NewInterpreter(input, output, parser)
		_ = bfm.Run()

		if output.String() != tt.wanted {
			t.Errorf("wrong output, want %s,got %s", tt.wanted, output.String())
		}
	}
}
