package cmd

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
	"github.com/Gealber/devto-cli/display"
)

type ProfileImageCommand Command

func NewProfileImageCmd() *ProfileImageCommand {
	return &ProfileImageCommand{
		Name:        "profile_images",
		Description: "Retrieve, Create and Update the profile image",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve profile image",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *ProfileImageCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		err = c.retrieveProfileImage()
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *ProfileImageCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *ProfileImageCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *ProfileImageCommand) SetData(data string) {
	c.Data = data
}

//retrieveProfileImage ...
func (c *ProfileImageCommand) retrieveProfileImage() CommandValidationError {
	ctx := context.Background()
	imgPro, err := api.RetrieveProfileImage(ctx, c.Data)
	if err != nil {
		return err
	}
	display.ProfileImageResponse(imgPro)
	return nil
}

//ActivateSubcommand ...
func (c *ProfileImageCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
