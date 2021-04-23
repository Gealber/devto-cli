package cmd

import (
	"github.com/Gealber/devto-cli/api"
	"text/tabwriter"
)

type AdminConfig Command

func NewAdminConfigCmd() *AdminConfig {
	return &AdminConfig{
		Name:        "admin-config",
		Description: "Retrieve or update admin configuration from dev.to",
		Subcommands: map[string]*Subcommand{
			"update": {
				Description: "Update admin configuration from dev.to",
				Active:      false,
			},
			"retrieve": {
				Description: "Retrieve admin configuration from dev.to",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *AdminConfig) Run() (*CommandResponse, CommandValidationError) {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	if c.Subcommands["retrieve"].Active {
		c.retrieve()
	} else if c.Subcommands["udpate"].Active {
		c.update()
	}
	return nil, nil
}

//Validate execute the command
func (c *AdminConfig) Validate() CommandValidationError {
	return nil
}

//Helper display info about the command
func (c *AdminConfig) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//retrieve ...
func (c *AdminConfig) retrieve() (*CommandResponse, CommandValidationError) {
	api.RetrieveAdminConfig()
	return nil, nil
}

//update ...
func (c *AdminConfig) update() (*CommandResponse, CommandValidationError) {
	api.UpdateAdminConfig()
	return nil, nil
}

//SetData ...
func (c *AdminConfig) SetData(data string) {
	c.Data = data
}

//ActivateSubcommand ...
func (c *AdminConfig) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError("Subcommand not found")
	}
	c.Subcommands[name].Active = true
	return nil
}
