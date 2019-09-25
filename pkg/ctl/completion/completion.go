package completion

import (
	"os"

	"github.com/kris-nova/logger"
	"github.com/spf13/cobra"
)

func Command(rootCmd *cobra.Command) *cobra.Command {
	var bashCompletionCmd = &cobra.Command{
		Use:   "bash",
		Short: "Generates bash completion scripts",
		Long: `To load completion run

. <(pulsarctl completion bash)

To configure your bash shell to load completions for each session add to your bashrc

# ~/.bashrc or ~/.profile
. <(pulsarctl completion bash)

If you are stuck on Bash 3 (macOS) use

source /dev/stdin <<<"$(pulsarctl completion bash)"
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenBashCompletion(os.Stdout)
		},
	}

	var zshCompletionCmd = &cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		Long: `To configure your zsh shell, run:

mkdir -p ~/.zsh/completion/
pulsarctl completion zsh > ~/.zsh/completion/_pulsarctl

and put the following in ~/.zshrc:

fpath=($fpath ~/.zsh/completion)
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenZshCompletion(os.Stdout)
		},
	}

	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generates shell completion scripts",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				logger.Debug("ignoring error %q", err.Error())
			}
		},
	}

	cmd.AddCommand(bashCompletionCmd)
	cmd.AddCommand(zshCompletionCmd)

	return cmd
}
