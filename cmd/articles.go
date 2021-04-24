package cmd

import (
	"fmt"
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
			"update": {
				Description: "Update and article",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *ArticlesCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		err = c.retrieve()
		if err != nil {
			return err
		}
	} else if c.Subcommands["update"].Active {
		article, err := processUpdate()
		if err != nil {
			return err
		}
		err = c.update(article)
		if err != nil {
			return err
		}
	} else if c.Subcommands["create"].Active {
		err = c.create()
		if err != nil {
			return err
		}
	}
	return nil
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
func (c *ArticlesCommand) retrieve() CommandValidationError {
	_, err := api.RetrieveArticles(c.Data)
	if err != nil {
		return err
	}
	return nil
}

//update ...
func (c *ArticlesCommand) update(data *api.ArticleEdit) CommandValidationError {
	_, err := api.UpdateArticle(c.Data, data)
	if err != nil {
		return err
	}
	return nil
}

//processUpdate read the data from the User input and put
//that data inside an ArticleEdit structure
func processUpdate() (*api.ArticleEdit, error) {
	//to store field from ArticleEdit
	article := new(api.ArticleEdit)
	err := processInput(article)
	if err != nil {
		return nil, err
	}
	return article, nil
}

//create ...
func (c *ArticlesCommand) create() CommandValidationError {
	return nil
}

//ActivateSubcommand ...
func (c *ArticlesCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
