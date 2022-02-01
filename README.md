# BrainFuck interpreter




# Description
Brainfuck interpeter is an interpeter for brainfuck language which is based on the principle of abstract structure tree.
the entry point of the code is *io.Reader*
supported operators that are initially enabled :

- *>* &nbsp;: &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;increments (increases by 1) the pointer.  
- *<* :&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; decrements (decreases by 1) the pointer.
- *+* &nbsp;:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; increments the byte of the array on which the pointer is positioned (the pointed byte).
- *-* &nbsp;&nbsp;:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; decrements the byte of the array on which the pointer is positioned (the pointed byte).
- *.* &nbsp;&nbsp;: &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;output of pointed byte (ASCII value).
- *,* &nbsp;&nbsp;: &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;entry of a byte in the table where the pointer is positioned (ASCII value).
- [ &nbsp;&nbsp;: &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;jumps to the instruction after the corresponding ] if the pointed byte is at 0.
- ] &nbsp;&nbsp;: &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;returns to the instruction after the [ if the pointed byte is different from 0.
  
supported operators that are initially disabled :
- ***²*** : calculate the square of the byte of the array on which the pointer is positioned (the pointed byte).

## refrences :
the book : writing an INTERPRETER in go https://edu.anarcho-copy.org/Programming%20Languages/Go/writing%20an%20INTERPRETER%20in%20go.pdf


## to run unit tests
```bash
go test
```
## to run component_tests
```bash
go install github.com/cucumber/godog/cmd/godog@v0.12.0
cd component_tests
godog
```
## to use 
- firstly you have to create an instance of a parser using io.Reader input 
```bashs
   parser := NewParser(code)
```
- declare your input of type *bytes.Buffer(the buffer that will be used for the table) and the output of type *bytes.Buffer
```bashs
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)
```
- then create the instance of your brainfuck (machine) 
```bashs
    bfm := NewInterpreter(input, output,parser)
```
this interpreter only to initialize the machine 
- You can now run your interpter using run() method
```bashs
    _ = bfm.Run()

```
## full example  
 ```bashs

 	import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	bf "github.com/Oussama1920/Brainfuck"
	)
	func main() {

	HelloWordInBrainFuck := `>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+`

	code := strings.NewReader(HelloWordInBrainFuck)
	parser := bf.NewParser(code)
	input := new(bytes.Buffer)
	output := new(bytes.Buffer)
	bfm := bf.NewInterpreter(input, output, parser)
	_ = bfm.Run()
	fmt.Println("code successfully compiled and output is : %s", output.String())
	}
```

## Add or Delete operators before interpreting 
the operators that can be activated/disactivated are : [+],[,], [-],[>],[<],[.] 
- to desactivate(delete) an operator before running the machine
  example: 
```bash
	code := strings.NewReader(HelloWordInBrainFuck)
	parser := NewParser(code)
	parser.Desactivate("+")

```
- to activate(add) an operator before running the machine
  example: 
```bash
	code := strings.NewReader(HelloWordInBrainFuck)
	parser := NewParser(code)
	parser.activate("²")
```
PLEASE NOTE :
- if an operator has been desactivated so the interpreter will ignore it.
- as the loop in tha main function of the brainfuck so we chose to DO NOT SUPPORT he desactivation so ```bash parser.Desactivate("[") ```
 or  ```bash
parser.Desactivate("]")```
 will not change anything :) 
- be careful !! we don't support any caracter outside of the list already mentioned ==> your code should obligatory only contains valid operators (Maybe in next version we will support that by deleting things out of tune)
   

