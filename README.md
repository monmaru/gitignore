# gitignore
Interactive CLI tool to generate .gitignore from [github/gitignore](https://github.com/github/gitignore).

![screenshot](screenshot.png)

## Installation
```
go get github.com/monmaru/gitignore
```

## Usage
```
NAME:
   gitignore - A new cli application

USAGE:
   gitignore [global options] command [command options] [arguments...]

VERSION:
   1.0

DESCRIPTION:
   generate a .gitignore from github/gitignore on GitHub

AUTHOR:
   monmaru

COMMANDS:
     list     List a collection of .gitignore templates
     dump     Dump a .gitignore to console
     new      Create a .gitignore to current directory
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## License
Licensed under the [MIT](LICENSE) License.