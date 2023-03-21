# got
![got](got.png)

A CLI written in GO for with some shortcuts for Git that I use daily.

## Motivation

I usually add a set of alias for some git commands that I use the most. The alias, are usually for commands that either I find long to type, or hard to remember because I don't use it that ofter.
e.g.
```shell
alias push='git push origin head'
alias stat='git status -s'
alias nb='git checkout -b'
alias goto='git checkout'
alias reset='git reset --hard'
alias ci='git commit -am'
alias gbr="git branch | grep -v "main" | xargs git branch -D"
```

Having it as shell script is fine, but inspired by [gut](https://github.com/gut-hub/gut), I decided to port my script to a CLI application

## What commands are available?

I'm adding the commands as I have some free time. So far, the following commands are available:

| got      | git           | description                                   |
|----------|---------------|-----------------------------------------------|
| got stat | git status -s | get a summary of the status of the local repo |


