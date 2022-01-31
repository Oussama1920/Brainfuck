# BrainFuck interpreter




# Description
Brainfuck interpeter is an interpeter for brainfuck language which is based on the principle of abstract structure tree.
the entry point of the code is *io.Reader*
supported operators that are initially enabled :

- *>* : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;increments (increases by 1) the pointer.  
- *<* :&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; decrements (decreases by 1) the pointer.
- *+* :&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; increments the byte of the array on which the pointer is positioned (the pointed byte).
- *-* :&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; decrements the byte of the array on which the pointer is positioned (the pointed byte).
- *.* : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;output of pointed byte (ASCII value).
- *,* : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;entry of a byte in the table where the pointer is positioned (ASCII value).
- [ : &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;jumps to the instruction after the corresponding ] if the pointed byte is at 0.
  
supported operators that are initially disabled :
- ***²*** : calculate the square of the byte of the array on which the pointer is positioned (the pointed byte).

## refrences :
the book : writing an INTERPRETER in go https://edu.anarcho-copy.org/Programming%20Languages/Go/writing%20an%20INTERPRETER%20in%20go.pdf


## to run unit tests
```bash
go test
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
HelloWordInBrainFuck := `>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+`

code := strings.NewReader(HelloWordInBrainFuck)
parser := NewParser(code)
input := new(bytes.Buffer)
output := new(bytes.Buffer)
bfm := NewInterpreter(input, output, parser)
_ = bfm.Run()
 fmt.Println("code successfully compiled and output is : %s", output.String())
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
NB :
- if an operator has been desactivated so the interpreter will ignore it.
- as the loop in tha main function of the brainfuck so we chose to DO NOT SUPPORT he desactivation so ```bash parser.Desactivate("[") ```
 or  ```bash
parser.Desactivate("]")```
 will not change anything :) 

