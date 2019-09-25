package cmdutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

const IncompatibleFlags = "cannot be used at the same time"

// NewVerbCmd defines a standard resource command
func NewResourceCmd(use, short, long string, aliases ...string) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Aliases: aliases,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}
}

var CheckNameArgError = defaultNameArgsError

var defaultNameArgsError = func(err error) {
	os.Exit(1)
}

// GetNameArg tests to ensure there is only 1 name argument
func GetNameArg(args []string) (string, error) {
	if len(args) > 1 || len(args) == 0 {
		logger.Critical("only one argument is allowed to be used as a name")
		err := errors.New("only one argument is allowed to be used as a name")
		CheckNameArgError(err)
		return "", err
	}
	if len(args) == 1 {
		return strings.TrimSpace(args[0]), nil
	}
	return "", nil
}

func GetNameArgs(args []string, check func(args []string) error) ([]string, error) {
	err := check(args)
	if err != nil {
		logger.Critical(err.Error())
		CheckNameArgError(err)
		//for testing
		return nil, err
	}
	return args, nil
}

func NewPulsarClient() pulsar.Client {
	return PulsarCtlConfig.Client(pulsar.V2)
}

func NewPulsarClientWithAPIVersion(version pulsar.APIVersion) pulsar.Client {
	return PulsarCtlConfig.Client(version)
}

func PrintJSON(w io.Writer, obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		_, _ = fmt.Fprintf(w, "unexpected response type: %v\n", err)
		return
	}
	_, _ = fmt.Fprintln(w, string(b))
}

func PrintError(w io.Writer, err error) {
	msg := err.Error()
	if pulsar.IsAdminError(err) {
		ae, _ := err.(pulsar.Error)
		msg = ae.Reason
	}
	fmt.Fprintln(w, "error:", msg)
}
