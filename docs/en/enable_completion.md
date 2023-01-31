<!--

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.

-->

## Enabling shell autocompletion

pulsarctl provides autocompletion support for Bash, Zsh, and Fish, which can save you a lot of typing.

### Zsh

The pulsarctl completion script for Zsh can be generated with the command `pulsarctl completion zsh`. Sourcing the completion script in your shell enables pulsarctl autocompletion.

To configure your zsh shell, run:

```bash
mkdir -p ~/.zsh/completion/
pulsarctl completion zsh > ~/.zsh/completion/_pulsarctl
```

Include the directory in your $fpath, for example by adding in ~/.zshrc:

```bash
fpath=($fpath ~/.zsh/completion)
source ~/.zshrc
```

You may have to force rebuild zcompdump:

```bash
rm -f ~/.zcompdump; compinit
```

### Bash

#### Introduction

The pulsarctl completion script for Bash can be generated with `pulsarctl completion bash`. Sourcing this script in your shell enables pulsarctl completion.

However, the pulsarctl completion script depends on `bash-completion` which you thus have to previously install.

> Warning: there are two versions of bash-completion, v1 and v2. V1 is for Bash 3.2 (which is the default on macOS), and v2 is for Bash 4.1+. The pulsarctl completion script doesn’t work correctly with bash-completion v1 and Bash 3.2. It requires bash-completion v2 and Bash 4.1+. Thus, to be able to correctly use pulsarctl completion on macOS, you have to install and use Bash 4.1+ (instructions). The following instructions assume that you use Bash 4.1+ (that is, any Bash version of 4.1 or newer).

#### Install bash-completion

> Note: As mentioned, these instructions assume you use Bash 4.1+, which means you will install bash-completion v2 (in contrast to Bash 3.2 and bash-completion v1, in which case pulsarctl completion won’t work).

You can test if you have bash-completion v2 already installed with `brew list | grep bash`. If not, you can install it with Homebrew:

```bash
brew install bash-completion@2
```

As stated in the output of this command, add the following to your ~/.bashrc file:

```bash
export BASH_COMPLETION_COMPAT_DIR="/usr/local/etc/bash_completion.d"
[[ -r "/usr/local/etc/profile.d/bash_completion.sh" ]] && . "/usr/local/etc/profile.d/bash_completion.sh"
```

#### Enable pulsarctl autocompletion

You now have to ensure that the pulsarctl completion script gets sourced in all your shell sessions. There are multiple ways to achieve this:

- First, you can use `bash` into the bash shell.

> Note: If you are using the bash shell, you can ignore it

- Add the completion script to the `/usr/local/etc/bash_completion.d` directory:

```bash
pulsarctl completion bash >/usr/local/etc/bash_completion.d/pulsarctl.bash
```

- Source the completion script in your `~/.bashrc` file:

```bash
echo 'source /usr/local/etc/bash_completion.d/pulsarctl.bash' >> ~/.bashrc
source ~/.bashrc
```

### Fish

To load completions once in your current session run:
		
```bash
pulsarctl completion fish | source
```

To load completions for each session, run:

```bash
pulsarctl completion fish > ~/.config/fish/completions/pulsarctl.fish
```
