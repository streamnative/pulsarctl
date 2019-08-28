package pulsar

var SPACES = "    "
var USED_FOR  = "USED FOR:"
var PERMISSION = "REQUIRED PERMISSION:"
var EXAMPLES = "EXAMPLES:"
var OUTPUT = "OUTPUT:"

type LongDescription struct {
	CommandUsedFor string
	CommandPermission string
	CommandExamples []Example
	CommandOutput string
}

type Example struct {
	Desc string
	Command string
}

func (desc *LongDescription) ToString() string {
	return USED_FOR + "\n" +
		   SPACES + desc.CommandUsedFor + "\n" +
		   PERMISSION + "\n" +
		   SPACES + desc.CommandPermission + "\n" +
		   EXAMPLES + "\n" +
		   desc.exampleToString(SPACES) +
		   OUTPUT + "\n" +
		   SPACES + desc.CommandOutput  + "\n"
}

func (desc *LongDescription) exampleToString(space string) string {
	var result string
	for _, v := range desc.CommandExamples {
		result += space + "#" + v.Desc + "\n" +
			      space + v.Command + "\n"
	}
	return result
}
