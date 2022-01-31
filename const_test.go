package brainfuck

// source : https://therenegadecoder.com/code/hello-world-in-brainfuck/
var HelloWordInBrainFuck = `>++++++++[<+++++++++>-]<.>++++[<+++++++>-]<+.+++++++..+++.>>++++++[<+++++++>-]<++.------------.>++++++[<+++++++++>-]<+.<.+++.------.--------.>>>++++[<++++++++>-]<+`

// source : http://esoteric.sange.fi/brainfuck/bf-source/prog/fibonacci.txt
var FibonacciInBrainfFuck = `+++++++++++>+>>>>++++++++++++++++++++++++++++++++++++++++++++>++++++++++++++++++++++++++++++++<<<<<<[>[>>>>>>+>+<<<<<<<-]>>>>>>>[<<<<<<<+>>>>>>>-]<[>++++++++++[-<-[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<[>>>+<<<-]>>[-]]<<]>>>[>>+>+<<<-]>>>[<<<+>>>-]+<[>[-]<[-]]>[<<+>>[-]]<<<<<<<]>>>>>[++++++++++++++++++++++++++++++++++++++++++++++++.[-]]++++++++++<[->-<]>++++++++++++++++++++++++++++++++++++++++++++++++.[-]<<<<<<<<<<<<[>>>+>+<<<<-]>>>>[<<<<+>>>>-]<-[>>.>.<<<[-]]<<[>>+>+<<<-]>>>[<<<+>>>-]<<[<+>-]>[<+>-]<<<-]`

// source : http://progopedia.com/example/factorial/18/
var FactorialInBrainFuck = `+++++++++++++++++++++++++++++++++>+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++>++++++++++>+++++++>>+<<[>++++++++++++++++++++++++++++++++++++++++++++++++.------------------------------------------------<<<<.-.>.<.+>>>>>>>++++++++++<<[->+>-[>+>>]>[+[-<+>]>+>>]<<<<<<]>[<+>-]>[-]>>>++++++++++<[->-[>+>>]>[+[-<+>]>+>>]<<<<<]>[-]>>[++++++++++++++++++++++++++++++++++++++++++++++++.[-]]<[++++++++++++++++++++++++++++++++++++++++++++++++.[-]]<<<++++++++++++++++++++++++++++++++++++++++++++++++.[-]<<<<<<.>>+>[>>+<<-]>>[<<<[>+>+<<-]>>[<<+>>-]>-]<<<<-]`

var testParser = []*Instruction{
	{additionalData: 5, id: Identifier{Token: op_inc_val, Value: "+"}},
	{additionalData: 2, id: Identifier{Token: op_dec_val, Value: "-"}},
	{additionalData: 4, id: Identifier{Token: op_jmp_fwd, Value: "["}},
	{additionalData: 1, id: Identifier{Token: op_dec_val, Value: "-"}},
	{additionalData: 2, id: Identifier{Token: op_jmp_bck, Value: "]"}},
}

//                                               -[--[+]--]
var testParseInnerLop = []*Instruction{
	{id: Identifier{Token: op_dec_val, Value: "-"}, additionalData: 1},
	{id: Identifier{Token: op_jmp_fwd, Value: "["}, additionalData: 7},
	{id: Identifier{Token: op_dec_val, Value: "-"}, additionalData: 2},
	{id: Identifier{Token: op_jmp_fwd, Value: "["}, additionalData: 5},
	{id: Identifier{Token: op_inc_val, Value: "+"}, additionalData: 1},
	{id: Identifier{Token: op_jmp_bck, Value: "]"}, additionalData: 3},
	{id: Identifier{Token: op_dec_val, Value: "-"}, additionalData: 2},
	{id: Identifier{Token: op_jmp_bck, Value: "]"}, additionalData: 1},
}
var testParseCursorMoving = []*Instruction{
	{id: Identifier{Token: op_inc_val, Value: "+"}, additionalData: 1},
	{id: Identifier{Token: op_inc_dp, Value: ">"}, additionalData: 3},
	{id: Identifier{Token: op_inc_val, Value: "+"}, additionalData: 7},
	{id: Identifier{Token: op_inc_dp, Value: ">"}, additionalData: 2},
	{id: Identifier{Token: op_inc_val, Value: "+"}, additionalData: 3},
	{id: Identifier{Token: op_dec_val, Value: "-"}, additionalData: 2},
	{id: Identifier{Token: op_dec_dp, Value: "<"}, additionalData: 1},
}
