package cmdutils

import (
	"encoding/json"
	"fmt"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"io"
	"os"
	"strings"
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
	if len(args) > 1 {
		logger.Critical("only one argument is allowed to be used as a name")
		os.Exit(1)
	}
	if len(args) == 1 {
		return strings.TrimSpace(args[0])
	}
	return ""
}

func NewPulsarClient() pulsar.Client {
	return PulsarCtlConfig.Client()
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