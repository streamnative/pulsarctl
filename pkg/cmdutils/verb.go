package cmdutils

import (
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"os"
)

// VerbCmd holds attributes that most of the commands use
type VerbCmd struct {
	Command      *cobra.Command
	FlagSetGroup *NamedFlagSetGroup
	NameArg      string
	NameArgs     []string
}

// AddVerbCmd create a registers a new command under the given resource command
func AddVerbCmd(flagGrouping *FlagGrouping, parentResourceCmd *cobra.Command, newVerbCmd func(*VerbCmd)) {
	verb := &VerbCmd{
		Command: &cobra.Command{},
	}
	verb.FlagSetGroup = flagGrouping.New(verb.Command)
	newVerbCmd(verb)
	verb.FlagSetGroup.AddTo(verb.Command)
	parentResourceCmd.AddCommand(verb.Command)
}

// SetDescription sets usage along with short and long descriptions as well as aliases
func (vc *VerbCmd) SetDescription(use, short, long string, aliases ...string) {
	vc.Command.Use = use
	vc.Command.Short = short
	vc.Command.Long = long
	vc.Command.Aliases = aliases
}

// SetRunFunc registers a command function
func (vc *VerbCmd) SetRunFunc(cmd func() error) {
	vc.Command.Run = func(_ *cobra.Command, _ []string) {
		run(cmd)
	}
}

// SetRunFuncWithNameArg registers a command function with an optional name argument
func (vc *VerbCmd) SetRunFuncWithNameArg(cmd func() error) {
	vc.Command.Run = func(_ *cobra.Command, args []string) {
		vc.NameArg = GetNameArg(args)
		run(cmd)
	}
}

func (vc *VerbCmd) SetRunFuncWithNameArgs(cmd func() error, checkArgs func(args []string) error) {
	vc.Command.Run = func(_ *cobra.Command, args []string) {
		vc.NameArgs = GetNameArgs(args, checkArgs)
		run(cmd)
	}
}

var ExecErrorHandler = defaultExecErrorHandler

var defaultExecErrorHandler = func(err error) {
	logger.Critical("%s\n", err.Error())
	os.Exit(1)
}

func run(cmd func() error) {
	if err := cmd(); err != nil {
		ExecErrorHandler(err)
	}
}
