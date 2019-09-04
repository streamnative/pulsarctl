package cluster

import (
	"bytes"
	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"os"
)

func TestClusterCommands(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer, execErr, nameErr, err error) {
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
	rootCmd.SetArgs(append([]string{"clusters"}, args...))

	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")
	flagGrouping := cmdutils.NewGrouping()
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.AddCommand(resourceCmd)
	err = rootCmd.Execute()

	return buf, execError, nameError, err
}

var (
	flag     bool
	basePath string
)

func TestTlsHelp(newVerb func(cmd *cmdutils.VerbCmd), args []string) (out *bytes.Buffer, err error) {
	var rootCmd = &cobra.Command{
		Use:   "pulsarctl [command]",
		Short: "a CLI for Apache Pulsar",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}

	if !flag {
		basePath, err = os.Getwd()
		if err != nil {
			return nil, err
		}

		err = os.Chdir("../../../")
		if err != nil {
			return nil, err
		}

		basePath, err = os.Getwd()
		if err != nil {
			return nil, err
		}
		flag = true
	}

	baseArgs := []string{
		"--auth-params",
		"{\"tlsCertFile\":\"" + basePath + "/test/auth/certs/client-cert.pem\"" +
			",\"tlsKeyFile\":\"" + basePath + "/test/auth/certs/client-key.pem\"}",
		"--tls-trust-cert-pat", basePath + "/test/auth/certs/cacert.pem",
		"--admin-service-url", "https://localhost:8443",
		"--tls-allow-insecure"}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs(append(baseArgs, args...))
	resourceCmd := cmdutils.NewResourceCmd(
		"clusters",
		"Operations about cluster(s)",
		"",
		"cluster")

	flagGrouping := cmdutils.NewGrouping()

	rootCmd.AddCommand(resourceCmd)
	cmdutils.AddVerbCmd(flagGrouping, resourceCmd, newVerb)
	rootCmd.PersistentFlags().AddFlagSet(cmdutils.PulsarCtlConfig.FlagSet())
	err = rootCmd.Execute()
	return buf, err
}
