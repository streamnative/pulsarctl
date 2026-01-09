// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

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
			return rootCmd.GenBashCompletionV2(os.Stdout, true)
		},
	}

	var zshCompletionCmd = &cobra.Command{
		Use:   "zsh",
		Short: "Generates zsh completion scripts",
		Long: `To configure your zsh shell, run:

mkdir -p ~/.zsh/completion/
pulsarctl completion zsh > ~/.zsh/completion/_pulsarctl

# Include the directory in your $fpath, for example by adding in ~/.zshrc:
fpath=($fpath ~/.zsh/completion)

# You may have to force rebuild zcompdump:
rm -f ~/.zcompdump; compinit
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenZshCompletion(os.Stdout)
		},
	}

	var fishCompletionCmd = &cobra.Command{
		Use:   "fish",
		Short: "Generates Fish completion scripts",
		Long: `To load completions run:

pulsarctl completion fish | source

		To load completions for each session, run:

pulsarctl completion fish > ~/.config/fish/completions/pulsarctl.fish
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCmd.GenFishCompletion(os.Stdout, true)
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
	cmd.AddCommand(fishCompletionCmd)

	return cmd
}
