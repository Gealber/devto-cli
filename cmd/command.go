package cmd

import (
	"errors"
	"fmt"
	"text/tabwriter"
)

//Command main struture for a command
type Command struct {
	Name        string
	Description string
	Data        string
	Response    *CommandResponse
	Subcommands map[string]*Subcommand
}

type Subcommand struct {
	Description string
	Active      bool
}

func Helper(name, description string, tw *tabwriter.Writer) {
	format := "\t\033[1;32m%v:\033[0m\t\033[3m%v\033[0m\n"
	fmt.Fprintf(tw, format, name, description)
}

//CommandResponse response by the api to the execution
type CommandResponse struct {
	Code        int
	Description string
}

//CommandValidationError ...
type CommandValidationError error

//NewCommandError ...
func NewCommandError(description string) CommandValidationError {
	return errors.New(description)
}

//CommandI actions that can be excecuted by a Command
type CommandI interface {
	Run() (*CommandResponse, CommandValidationError)
	Helper(*tabwriter.Writer)
}

//ActivateSubcommand ...
func (c *Command) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError("Subcommand not found")
	}
	c.Subcommands[name].Active = true
	return nil
}
