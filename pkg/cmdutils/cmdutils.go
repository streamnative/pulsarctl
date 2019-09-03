package cmdutils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

const IncompatibleFlags = "cannot be used at the same time"

// NewVerbCmd defines a standard resource command
func NewResourceCmd(use, short, long string, aliases ...string) *cobra.Command {
	return &cobra.Command{
		Use:		use,
		Short: 		short,
		Long:		long,
		Aliases:	aliases,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}
}

// GetNameArg tests to ensure there is only 1 name argument
func GetNameArg(args []string) string {
	if len(args) > 1 || len(args) == 0 {
		logger.Critical("only one argument is allowed to be used as a name")
		os.Exit(1)
	}
	if len(args) == 1 {
		return strings.TrimSpace(args[0])
	}
	return ""
}

func GetNameArgs(args []string, check func(args []string) error) []string {
	err := check(args)
	if err != nil {
		logger.Critical(err.Error())
		os.Exit(1)
	}
	return args
}

func NewPulsarClient() pulsar.Client {
	return PulsarCtlConfig.Client(pulsar.V2)
}

func NewPulsarClientWithApiVersion(version pulsar.ApiVersion) pulsar.Client {
	return PulsarCtlConfig.Client(version)
}

func PrintJson(w io.Writer, obj interface{}) {
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "unexpected response type: %v\n", err)
		return
	}
	fmt.Fprintln(w, string(b))
}


func PrintError(w io.Writer, err error) {
	msg := err.Error()
	if pulsar.IsAdminError(err) {
		ae, _ := err.(pulsar.Error)
		msg = ae.Reason
	}
	fmt.Fprintln(w, "error:", msg)
}