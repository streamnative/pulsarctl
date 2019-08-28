package pulsar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongDescription_exampleToString(t *testing.T)  {
	desc := LongDescription{}
	example := Example{
		Desc: "command description",
		Command: "command",
	}
	desc.CommandExamples = []Example{example}
	res := desc.exampleToString("    ")

	expect := "    #command description\n" +
		      "    command\n"

	assert.Equal(t, expect, res)
}

func TestLongDescription_ToString(t *testing.T) {
	desc := LongDescription{}
	desc.CommandUsedFor = "command used for"
	desc.CommandPermission = "command permission"
	example := Example{}
	example.Desc = "command description"
	example.Command = "command"
	desc.CommandExamples = []Example{example}
	desc.CommandOutput = "out"

	expect := "USED FOR:\n" +
		"    " + desc.CommandUsedFor + "\n" +
		"REQUIRED PERMISSION:\n" +
		"    " + desc.CommandPermission  + "\n" +
		"EXAMPLES:\n" +
		"    " + "#" + example.Desc + "\n" +
		"    " + example.Command + "\n" +
		"OUTPUT:\n" +
		"    " + desc.CommandOutput  + "\n"

	result := desc.ToString()

	assert.Equal(t, expect, result)
}