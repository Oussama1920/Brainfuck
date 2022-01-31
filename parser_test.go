package brainfuck

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	for _, tt := range []struct {
		name                      string
		wanted                    []*Instruction
		input                     string
		wantedNumberOfInstruction int
	}{
		{"nominal", testParser, "+++++--[-]", 5},
		{"inner loop", testParseInnerLop, "-[--[+]--]", 8},
		{"cursor moving", testParseCursorMoving, "+>>>+++++++>>+++--<", 7},
	} {
		input := strings.NewReader(tt.input)
		parser := NewParser(input)
		instructions := parser.Parse()
		if len(instructions) != tt.wantedNumberOfInstruction {
			t.Errorf("wrong length, expected, %v got %+v", tt.wantedNumberOfInstruction, len(instructions))
		}
		for i, v := range tt.wanted {
			if *v != *instructions[i] {
				t.Errorf("wrong instruction. expected %+v got %+v", *v, *instructions[i])
			}
		}
	}
}

func Test_Desactivate(t *testing.T) {
	for _, tt := range []struct {
		name                      string
		operatoTodesactivate      string
		input                     string
		wantedNumberOfInstruction int
	}{
		{"desactivate +", "+", "+>>>+++++++>>+++--<", 4},
		{"desactivate >", ">", "+>>>+++++++>>+++--<", 5},
		{"desactivate -", "-", "+>>>+++++++>>+++--<", 6},
	} {
		code := strings.NewReader(tt.input)
		parser := NewParser(code)
		parser.Desactivate(tt.operatoTodesactivate)
		instructions := parser.Parse()
		if len(instructions) != tt.wantedNumberOfInstruction {
			t.Errorf("wrong result, got %d, instead of %d", len(instructions), tt.wantedNumberOfInstruction)
		}
	}
}

func Test_Activate(t *testing.T) {
	for _, tt := range []struct {
		name                      string
		operatoTodesactivate      string
		input                     string
		wantedNumberOfInstruction int
	}{
		{"Activate ²", "²", "+>>>+++++++>>+++--<²", 8},
		{"Activate +", "+", "+>>>+++++++>>+++--<", 7},
	} {
		code := strings.NewReader(tt.input)
		parser := NewParser(code)
		parser.Activate(tt.operatoTodesactivate)
		instructions := parser.Parse()
		if len(instructions) != tt.wantedNumberOfInstruction {
			t.Errorf("wrong result, got %d, instead of %d", len(instructions), tt.wantedNumberOfInstruction)
		}
	}
}
