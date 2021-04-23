package cmd

import (
	"os"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
)

type ArticlesCommand Command

func NewArticlesCmd() *ArticlesCommand {
	return &ArticlesCommand{
		Name:        "articles",
		Description: "Retrieve, Create and Update the articles",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrive articles",
				Active:      false,
			},
			"create": {
				Description: "Create article",
				Active:      false,
			},
			"udpate": {
				Description: "Update and article",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *ArticlesCommand) Run() (*CommandResponse, CommandValidationError) {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	if c.Subcommands["retrieve"].Active {
		c.retrieve()
	} else if c.Subcommands["udpate"].Active {
		c.update()
	} else if c.Subcommands["create"].Active {
		c.create()
	}
	return nil, nil
}

//Validate check for the preconditions to execute this command
func (c *ArticlesCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *ArticlesCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *ArticlesCommand) SetData(data string) {
	c.Data = data
}

//retrieve ...
func (c *ArticlesCommand) retrieve(username string) (*CommandResponse, CommandValidationError) {
	api.RetrieveArticles(username)
	return nil, nil
}

//update ...
func (c *ArticlesCommand) update(id string, data *api.ArticleEdit) (*CommandResponse, CommandValidationError) {
	api.UpdateArticle(id, data)
	return nil, nil
}

//create ...
func (c *ArticlesCommand) create() (*CommandResponse, CommandValidationError) {
	return nil, nil
}

//ActivateSubcommand ...
func (c *ArticlesCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError("Subcommand not found")
	}
	c.Subcommands[name].Active = true
	return nil
}
