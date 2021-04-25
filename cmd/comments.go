package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
)

type CommentsCommand Command

func NewCommentsCommand() *CommentsCommand {
	return &CommentsCommand{
		Name:        "comments",
		Description: "Retrieve comments",
		Subcommands: map[string]*Subcommand{
			"retrieve_aid": {
				Description: "Retrieve comments by article id",
				Active:      false,
			},
			"retrieve_pid": {
				Description: "Retrieve comments by podcast id",
				Active:      false,
			},
			"retrieve_id": {
				Description: "Retrieve comments by comment id",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *CommentsCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve_aid"].Active {
		err = c.retrieveByAID()
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_pid"].Active {
		err = c.retrieveByPID()
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve_id"].Active {
		err = c.retrieveByID()
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *CommentsCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *CommentsCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *CommentsCommand) SetData(data string) {
	c.Data = data
}

//retrieveByAID ...
func (c *CommentsCommand) retrieveByAID() CommandValidationError {
	queries := &api.CommentQuery{
		AID: c.Data,
	}
	_, err := api.RetrieveComments(queries)
	if err != nil {
		return err
	}
	return nil
}

//retrieveByPID ...
func (c *CommentsCommand) retrieveByPID() CommandValidationError {
	queries := &api.CommentQuery{
		PID: c.Data,
	}
	_, err := api.RetrieveComments(queries)
	if err != nil {
		return err
	}
	return nil
}

//retrieveByID ...
func (c *CommentsCommand) retrieveByID() CommandValidationError {
	_, err := api.RetrieveComment(c.Data)
	if err != nil {
		return err
	}
	return nil
}

//ActivateSubcommand ...
func (c *CommentsCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
