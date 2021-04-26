package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
	"github.com/Gealber/devto-cli/display"
)

type FollowersCommand Command

func NewFollowersCommand() *FollowersCommand {
	return &FollowersCommand{
		Name:        "followers",
		Description: "Retrieve followers",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve followers",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *FollowersCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		queries, err := processFollowersQueries()
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
func (c *FollowersCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *FollowersCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *FollowersCommand) SetData(data string) {
	c.Data = data
}

//retrieve ...
func (c *FollowersCommand) retrieve(queries *api.FollowersQuery) CommandValidationError {
	followers, err := api.RetrieveFollowers(queries)
	if err != nil {
		return err
	}
	display.FollowersResponse(followers)
	return nil
}

//processQueries read the data from the User input and put
//that data inside an GetArticleQuery structure
func processFollowersQueries() (*api.FollowersQuery, error) {
	//to store field from GetArticleQuery
	queries := new(api.FollowersQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//ActivateSubcommand ...
func (c *FollowersCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
