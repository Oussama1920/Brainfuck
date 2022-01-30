package brainfuck

import (
	"strings"
	"testing"
)

func TestScanner_Read(t *testing.T) {
	r := strings.NewReader("<>   this is string [+++]")
	s := NewScanner(r)
	token := s.Scan()

	if token.Value != "<" {
		t.Errorf("expect < given %q", token.Value)
	}
}

func TestScanner_Scan(t *testing.T) {
	//below string contains long white space and three runes
	r := strings.NewReader("[+]")
	s := NewScanner(r)

	// read  [ rune
	token := s.Scan()

	if token.Value != "[" {
		t.Errorf("expect [ given %q", token.Value)
	}

	// read  + rune
	token = s.Scan()
	if token.Value != "+" {
		t.Errorf("expect + given %q", token.Value)
	}

	// read the last rune
	token = s.Scan()
	if token.Value != "]" {
		t.Errorf("expect ] given %q", token.Value)
	}

}
