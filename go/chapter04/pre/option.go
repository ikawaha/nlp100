package main

import (
	"flag"
)

type Option struct {
	Input string
	Output string

	flagSet *flag.FlagSet
}

func NewOption() *Option {
	ret := &Option{
		flagSet: flag.NewFlagSet(commandName, flag.ContinueOnError),
	}
	ret.flagSet.StringVar(&ret.Input, "input", "", "input file (optional: default stdin)")
	ret.flagSet.StringVar(&ret.Output, "output", "", "output file (optional: default stdout)")
	return ret
}

func (o *Option) Parse(args []string) error {
	if err := o.flagSet.Parse(args); err != nil {
		return err
	}
	return nil
}

func (o *Option) PrintDefaults() {
	o.flagSet.PrintDefaults()
}

