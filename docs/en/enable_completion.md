## Enabling shell autocompletion

pulsarctl provides autocompletion support for Bash and Zsh, which can save you a lot of typing.

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
```

You may have to force rebuild zcompdump:

```bash
rm -f ~/.zcompdump; compinit
```

### Bash


