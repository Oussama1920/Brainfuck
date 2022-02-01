package component_tests

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	bf "github.com/Oussama1920/Brainfuck"
	"github.com/cucumber/godog"
)

var x = bf.BrainFuck{}

type testClient struct {
	brainFuck        *bf.BrainFuck
	output           *bytes.Buffer
	input            *bytes.Buffer
	parser           *bf.Parser
	instructionCount int
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	// Standards interface to io
	output := new(bytes.Buffer)
	input := new(bytes.Buffer)

	testClient := &testClient{
		brainFuck: &bf.BrainFuck{},
		output:    output,
		input:     input,
	}
	ctx.Step(`^The user is going to develop$`, testClient.theUserIsGoingToDevelop)
	ctx.Step(`^The user initialize the input with data "([^"]*)" to the interpreter$`, testClient.theUserInitializeTheData)
	ctx.Step(`^The user compile the code$`, testClient.theUserCompileTheCode)
	ctx.Step(`^The user activate the operator  "([^"]*)"$`, testClient.theUserActivateOperator)
	ctx.Step(`^The user desactivate the operator  "([^"]*)"$`, testClient.theUserDesactivateOperator)
	ctx.Step(`^The output should be "([^"]*)"$`, testClient.thereShouldBeOutput)
	ctx.Step(`^The number of instruction should be (\d+)$`, testClient.theNumberOfInstructionShouldBe)
}
func (c *testClient) theUserIsGoingToDevelop() error {
	return nil
}
func (c *testClient) theUserCompileTheCode() error {
	c.brainFuck = bf.NewInterpreter(c.input, c.output, c.parser)

	// Store the result in output interface
	_ = c.brainFuck.Run()

	return nil
}
func (c *testClient) theUserInitializeTheData(arg1 string) error {
	if strings.Compare(arg1, "") != 0 {
		data, err := ioutil.ReadFile(arg1)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		code := strings.NewReader(string(data))
		// initialize the Parser with input
		c.parser = bf.NewParser(code)

	}

	return nil
}
func (c *testClient) thereShouldBeOutput(arg1 string) error {
	if c.output.String() != arg1 {
		return fmt.Errorf("expected output to be %s , but there is %s", arg1, c.output.String())
	}
	return nil
}
func (c *testClient) theUserDesactivateOperator(arg1 string) error {
	c.parser.Desactivate(arg1)
	instructions := c.parser.Parse()
	c.instructionCount = len(instructions)
	return nil
}
func (c *testClient) theUserActivateOperator(arg1 string) error {
	c.parser.Activate(arg1)
	instructions := c.parser.Parse()
	c.instructionCount = len(instructions)
	return nil
}

func (c *testClient) theNumberOfInstructionShouldBe(arg1 int) error {
	if c.instructionCount != arg1 {
		return fmt.Errorf("expected output to be %d , but there is %d", arg1, c.instructionCount)
	}

	return nil
}
