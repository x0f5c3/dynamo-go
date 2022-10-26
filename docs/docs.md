# dynamo-go

## Usage
> Secure self-hosted dynamic DNS server

dynamo-go

## Flags
|Flag|Usage|
|----|-----|
|`-d, --debug`|enable debug messages|
|`--no-update`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`dynamo-go completion`|Generate the autocompletion script for the specified shell|
|`dynamo-go help`|Help about any command|
# ... completion
`dynamo-go completion`

## Usage
> Generate the autocompletion script for the specified shell

dynamo-go completion

## Description

```
Generate the autocompletion script for dynamo-go for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`dynamo-go completion bash`|Generate the autocompletion script for bash|
|`dynamo-go completion fish`|Generate the autocompletion script for fish|
|`dynamo-go completion powershell`|Generate the autocompletion script for powershell|
|`dynamo-go completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`dynamo-go completion bash`

## Usage
> Generate the autocompletion script for bash

dynamo-go completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(dynamo-go completion bash)

To load completions for every new session, execute once:

#### Linux:

	dynamo-go completion bash > /etc/bash_completion.d/dynamo-go

#### macOS:

	dynamo-go completion bash > /usr/local/etc/bash_completion.d/dynamo-go

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`dynamo-go completion fish`

## Usage
> Generate the autocompletion script for fish

dynamo-go completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	dynamo-go completion fish | source

To load completions for every new session, execute once:

	dynamo-go completion fish > ~/.config/fish/completions/dynamo-go.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`dynamo-go completion powershell`

## Usage
> Generate the autocompletion script for powershell

dynamo-go completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	dynamo-go completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`dynamo-go completion zsh`

## Usage
> Generate the autocompletion script for zsh

dynamo-go completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	dynamo-go completion zsh > "${fpath[1]}/_dynamo-go"

#### macOS:

	dynamo-go completion zsh > /usr/local/share/zsh/site-functions/_dynamo-go

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`dynamo-go help`

## Usage
> Help about any command

dynamo-go help [command]

## Description

```
Help provides help for any command in the application.
Simply type dynamo-go help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 26 October 2022**
