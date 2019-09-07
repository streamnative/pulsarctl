package test

import (
	"bytes"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func TestTopicCommands(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer, execErr, nameErr, err error) {
	var execError error
	cmdutils.ExecErrorHandler = func(err error) {
		execError = err
	}

	var nameError error
	cmdutils.CheckNameArgError = func(err error) {
		nameError = err
	}

	var rootCmd = &cobra.Command{
		Use:   "pulsarctl [command]",
		Short: "a CLI for Apache Pulsar",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(append([]string{"topics"}, args...))

	resourceCmd := cmdutils.NewResourceCmd(
		"topics",
		"Operations about topics(s)",
		"",
		"topic")
	flagGrouping := cmdutils.NewGrouping()
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.AddCommand(resourceCmd)
	err = rootCmd.Execute()

	return buf, execError, nameError, err
}
