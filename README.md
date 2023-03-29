[![.github/workflows/ci.yaml](https://github.com/flavio1110/got/actions/workflows/ci.yaml/badge.svg)](https://github.com/flavio1110/got/actions/workflows/ci.yaml)
# got
![got](got.png)

A CLI written in GO to extend the `git` cli adding some shortcuts and useful functions.
All regular commands used with `git` will work with `got`.

## Motivation

I usually add a set of alias for some git commands that I use the most. The alias, are usually for commands that either I find long to type, or hard to remember because I don't use it that ofter.
e.g.
```shell
alias push='git push origin head'
alias stat='git status -s'
alias gbr="git branch | grep -v "main" | xargs git branch -D"
```

Having it as shell script is fine, but inspired by [gut](https://github.com/gut-hub/gut), I decided to port my script to a CLI application

## Install

```
go install github.com/flavio1110/got
```

## What commands are available?

All commands from `git ` CLI, and:

| got            | git                                    | description                                                       |
|----------------|----------------------------------------|-------------------------------------------------------------------|
| got stat       | git status -s                          | Get a summary of the status of the local repo                     |
| got rmb        | ...                                    | Remove all branches except main, master, and the current one      |
| got nb name    | git checkout -b name                   | Create a new branch and switch to it                              |
| got push       | git push origin head                   | Push the current head to the remote origin                        |
| got ci message | git add --all && git commit -m message | Stage all modifications and commit them with the provided message |

I'm adding the commands as I have some free time. So far, the following commands are available:

### Scaffold a  new command
1 - Install the Cobra CLI Generator
More info <https://github.com/spf13/cobra-cli/blob/main/README.md>
```
 go install github.com/spf13/cobra-cli@latest
```
2 - Add the command
```
cobra-cli add [name] --author "Flavio Silva flavio1110@gmaill.com"
```