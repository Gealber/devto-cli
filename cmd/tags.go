package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
	"github.com/Gealber/devto-cli/display"
)

type TagsCommand Command

func NewTagsCommand() *TagsCommand {
	return &TagsCommand{
		Name:        "tags",
		Description: "Retrieve tags that I follow",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve all tags",
				Active:      false,
			},
			"retrieve_follows": {
				Description: "Retrieve tags that I follow",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *TagsCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve_follows"].Active {
		err = c.retrieveFollows()
		if err != nil {
			return err
		}
	} else if c.Subcommands["retrieve"].Active {
		queries, err := processTagsQueries()
		if err != nil {
			return err
		}
		err = c.retrieve(queries)
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *TagsCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *TagsCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *TagsCommand) SetData(data string) {
	c.Data = data
}

//retrieveFollows ...
func (c *TagsCommand) retrieveFollows() CommandValidationError {
	tags, err := api.RetrieveTagsIFollow()
	if err != nil {
		return err
	}
	display.FollowTagsResponse(tags)
	return nil
}

//retrieve ...
func (c *TagsCommand) retrieve(queries *api.CommonQuery) CommandValidationError {
	tags, err := api.RetrieveTags(queries)
	if err != nil {
		return err
	}
	display.TagsResponse(tags)
	return nil
}

//processTagsQueries read the data from the User input and put
//that data inside an TagsQuery structure
func processTagsQueries() (*api.CommonQuery, error) {
	//to store field from CommonQuery
	queries := new(api.CommonQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//ActivateSubcommand ...
func (c *TagsCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
