package brainfuck

import (
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	input := strings.NewReader("+++++ -- [-]")
	p := NewParser(input)
	instructions := p.Parse()
	// since we are folding instructions
	// there is one instruction +, but 4 times
	if len(instructions) != 5 {
		t.Errorf("wrong length, expected 5 got %+v", len(instructions))
	}
	expected := []*Instruction{
		{c: 5, id: Identifier{token: op_inc_val, Value: "+"}},
		{c: 2, id: Identifier{token: op_dec_val, Value: "-"}},
		{c: 4, id: Identifier{token: op_jmp_fwd, Value: "["}},
		{c: 1, id: Identifier{token: op_dec_val, Value: "-"}},
		{c: 2, id: Identifier{token: op_jmp_bck, Value: "]"}},
	}
	for i, v := range expected {
		if *v != *instructions[i] {
			t.Errorf("wrong instruction. expected %+v got %+v", *v, *instructions[i])
		}
	}
}

func TestInnerLoops(t *testing.T) {
	input := strings.NewReader("-[--[+]--]")
	p := NewParser(input)
	instructions := p.Parse()
	expected := []*Instruction{
		{id: Identifier{token: op_dec_val, Value: "-"}, c: 1},
		{id: Identifier{token: op_jmp_fwd, Value: "["}, c: 7},
		{id: Identifier{token: op_dec_val, Value: "-"}, c: 2},
		{id: Identifier{token: op_jmp_fwd, Value: "["}, c: 5},
		{id: Identifier{token: op_inc_val, Value: "+"}, c: 1},
		{id: Identifier{token: op_jmp_bck, Value: "]"}, c: 3},
		{id: Identifier{token: op_dec_val, Value: "-"}, c: 2},
		{id: Identifier{token: op_jmp_bck, Value: "]"}, c: 1},
	}

	for i, v := range expected {
		if *v != *instructions[i] {
			t.Errorf("incorrect instruction. expected %+v got %+v", *v, *instructions[i])
		}
	}
}

func Test_MoveBetweenCells(t *testing.T) {
	input := strings.NewReader("+>>>+++++++>>+++ --<<")
	p := NewParser(input)
	instructions := p.Parse()
	expected := []*Instruction{
		{id: Identifier{token: op_inc_val, Value: "+"}, c: 1},
		{id: Identifier{token: op_inc_dp, Value: ">"}, c: 3},
		{id: Identifier{token: op_inc_val, Value: "+"}, c: 7},
		{id: Identifier{token: op_inc_dp, Value: ">"}, c: 2},
		{id: Identifier{token: op_inc_val, Value: "+"}, c: 3},
		{id: Identifier{token: op_dec_val, Value: "-"}, c: 2},
		{id: Identifier{token: op_dec_dp, Value: "<"}, c: 2},
	}

	for i, v := range expected {
		if *v != *instructions[i] {
			t.Errorf("incorrect instruction. expected %+v got %+v", *v, *instructions[i])
		}
	}

}
