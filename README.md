# 🧞Genie

A command-line tool to generate fake documents and data for developers.

## Contents
- [Requirements](#requirements)
- [Installation](#installation)
- [How to use it?](#how-to-use-it)
- [Examples](#examples)

## Requirements
- [Go 1.22+](https://go.dev/dl/)
- make (optional, for Makefile commands)
  - macOS/Linux: already installed
  - Windows: `choco install make`



## Installation
Installation guide for the CLI tool.
### Using Make (recommended)

```shell
# clone the repository
$ git clone https://github.com/leoandretta/genie-cli.git
$ cd genie


# install globally
$ make install
```
### Manual

```shell
git clone https://github.com/leoandretta/genie-cli.git
cd genie-cli
go build -o genie ./cmd/cli
```
> After installing, make sure $GOPATH/bin is in your $PATH.
> 
> macOS/Linux — add to `~/.zshrc` or `~/.bashrc`:
> ```shell
> export PATH="$PATH:$(go env GOPATH)/bin"
> ```
> Windows — add to your user environment variables:
> ```shell
> C:\Users\<YourUser>\go\bin
> ```


## How to use it

You can start by running the command with the `--help` flag
```shell
$ genie --help
Usage: genie <command> [flags]

Available commands:
  generate       Generates a valid value for the option requested

Try: genie <command> --help for help with each available command.
```
### Commands
| Command  | Description                    |
| -------- | ------------------------------ |
| generate | Generates a valid value        |

> Run `genie <command> --help` to see flags for a specific command.

#### Generate 
**Available Options**
- CPF (identification number for individuals in Brazil)
- CNPJ (identification number for businesses in Brazil)

**Flags**
| Flag |  Default   | Description                                    |
| ---- | ---------- | ---------------------------------------------- |
| `-f` | (required) | Returns the value formatted in standard format | 
> Run `genie generate <option> --help` to see flags for a specific option.


## Examples

```shell
# generate a valid cpf value
$ genie generate cpf

# generate a valid cpf value formatted in standart cpf format (###.###.###-##)
$ genie generate cpf -f

# generate a valid cnpj value
$ genie generate cnpj

# generate a valid cnpj value formatted in standard cnpj format (##.###.###/####-##)
$ genie generate cnpj -f
```