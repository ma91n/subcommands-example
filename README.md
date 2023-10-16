# subcommands-example

google/subcommands implements sample

## Usage

```sh
$ go build -o subclip . 
$ subclip help
Usage: subclip <flags> <subcommand> <subcommand args>

Subcommands:
commands         list all command names
flags            describe all known top-level flags
help             describe subcommands and their syntax
print            Print clipboard to stdout.
write            Write to clipboard


$ subclip help print
print [-n number] [-trim]:
  Print clipboard content.
  -n int
        display within particular line number
  -trim
        enable trimming space chars

$ subclip help write
write [text]:
  Write to clipboard.
```

