package component_tests

import (
	"bytes"
	"fmt"
	"strings"

	bf "github.com/Oussama1920/Brainfuck"
	"github.com/cucumber/godog"
)

var x = bf.BrainFuck{}

type testClient struct {
	brainFuck *bf.BrainFuck
	output    *bytes.Buffer
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	output := new(bytes.Buffer)
	// Standards interface to io
	testClient := &testClient{
		brainFuck: &bf.BrainFuck{},
		output:    output,
	}
	ctx.Step(`^The user is going to develop$`, testClient.theUserIsGoingToDevelop)
	ctx.Step(`^The user initialize the input with data "([^"]*)" to the interpreter$`, testClient.theUserInitializeTheData)
	ctx.Step(`^The output should be (\d+)$`, testClient.thereShouldBeOutput)

}
func (c *testClient) theUserIsGoingToDevelop() error {
	return nil
}
func (c *testClient) theUserInitializeTheData(arg1 string) error {
	code := strings.NewReader(arg1)
	// initialize the Parser with input
	parser := bf.NewParser(code)

	// Standards interface to io
	input := new(bytes.Buffer)
	// initialize the machine
	c.brainFuck = bf.NewInterpreter(input, c.output, parser)

	// Store the result in output interface
	_ = c.brainFuck.Run()
	return nil
}
func (c *testClient) thereShouldBeOutput(arg1 string) error {
	if c.output.String() != arg1 {
		return fmt.Errorf("expected %s godogs to be remaining, but there is %s", arg1, c.output.String())
	}
	return nil
}
