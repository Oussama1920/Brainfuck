package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/Azhovan/brainfuck"
)

func mainh() {
	// create new io.Reader from inputs
	/* code := strings.NewReader(`
	+++++++++++
	>+>>>>++++++++++++++++++++++++++++++++++++++++++++
	>++++++++++++++++++++++++++++++++<<<<<<[>[>>>>>>+>
	+<<<<<<<-]>>>>>>>[<<<<<<<+>>>>>>>-]<[>++++++++++[-
	<-[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<[>>>+<<<
	-]>>[-]]<<]>>>[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]
	>[<<+>>[-]]<<<<<<<]>>>>>[+++++++++++++++++++++++++
	+++++++++++++++++++++++.[-]]++++++++++<[->-<]>++++
	++++++++++++++++++++++++++++++++++++++++++++.[-]<<
	<<<<<<<<<<[>>>+>+<<<<-]>>>>[<<<<+>>>>-]<-[>>.>.<<<
	[-]]<<[>>+>+<<<-]>>>[<<<+>>>-]<<[<+>-]>[<+>-]<<<-]	`) */
	fmt.Print("Enter text: ")
	//	reader := bufio.NewReader(os.Stdin)

	// initialize the Parser with input
	parser := brainfuck.NewParser(os.Stdin)
	// Standards interface to io
	output := new(bytes.Buffer)

	// initialize the machine
	bfm := brainfuck.NewInterpreter(os.Stdin, output, parser)

	// Store the result in output interface
	_ = bfm.Run()

	// print the result
	fmt.Println(output.String())

}
