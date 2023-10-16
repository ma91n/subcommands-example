package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/google/subcommands"
	"golang.design/x/clipboard"
)

type printCmd struct {
	num  int
	trim bool
}

func (*printCmd) Name() string     { return "print" }
func (*printCmd) Synopsis() string { return "Print clipboard to stdout." }
func (*printCmd) Usage() string {
	return `print [-n number] [-trim]:
  Print clipboard content.
`
}

func (p *printCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.num, "n", 0, "display within particular line number")
	f.BoolVar(&p.trim, "trim", false, "enable trimming space chars")

	subcommands.ImportantFlag("n")
}

func (p *printCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	err := clipboard.Init()
	if err != nil {
		log.Printf("[clip] %v\n", err)
		return subcommands.ExitFailure
	}

	reader := bufio.NewReader(bytes.NewReader(clipboard.Read(clipboard.FmtText)))
	for i := 0; ; i++ {
		if p.num != 0 && i == p.num {
			break
		}
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Printf("[clip] %v\n", err)
			return subcommands.ExitFailure
		}

		out := string(line)
		if p.trim {
			out = strings.TrimSpace(out)
		}
		fmt.Println(out)
	}

	return subcommands.ExitSuccess
}

type writeCmd struct{}

func (*writeCmd) Name() string     { return "write" }
func (*writeCmd) Synopsis() string { return "Write to clipboard" }
func (*writeCmd) Usage() string {
	return `write [text]:
  Write to clipboard.
`
}

func (p *writeCmd) SetFlags(_ *flag.FlagSet) {}

func (p *writeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...any) subcommands.ExitStatus {
	err := clipboard.Init()
	if err != nil {
		log.Printf("[clip] %v\n", err)
		return subcommands.ExitFailure
	}
	clipboard.Write(clipboard.FmtText, []byte(strings.Join(f.Args(), "\n")))
	return subcommands.ExitSuccess
}
