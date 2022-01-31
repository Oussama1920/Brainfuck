package brainfuck

import (
	"strings"
	"testing"
)

func TestScanner(t *testing.T) {
	for _, tt := range []struct {
		name   string
		wanted string
		input  string
		length int
	}{
		{"test Read", "<>", "<>", 2},
		{"test Read", "<>++--", "<>++--", 6},
		{"test Read", "[+]", "[+]", 3},
	} {
		n := 0
		code := strings.NewReader(tt.input)
		scanner := NewScanner(code)
		for n < tt.length {
			token := scanner.Scan()
			if token.Value != string(tt.wanted[n]) {
				t.Errorf("wrong output, want %s,got %s", tt.wanted, token.Value)
			}
			n++
		}

	}
}
