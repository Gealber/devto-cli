package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
)

type ReadingListsCommand Command

func NewReadingListsCommand() *ReadingListsCommand {
	return &ReadingListsCommand{
		Name:        "reading_lists",
		Description: "Retrieve reading lists",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve reading lists",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *ReadingListsCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		queries, err := processReadingListsQueries()
		if err != nil {
			return err
		}
		err = c.retrieveReadingList(queries)
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *ReadingListsCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *ReadingListsCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *ReadingListsCommand) SetData(data string) {
	c.Data = data
}

//processReadingListsQueries read the data from the User input and put
//that data inside an ListingQuery structure
func processReadingListsQueries() (*api.ReadingListQuery, error) {
	//to store field from ReadingListEpisodesQuery
	queries := new(api.ReadingListQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//retrieveReadingList ...
func (c *ReadingListsCommand) retrieveReadingList(queries *api.ReadingListQuery) CommandValidationError {
	_, err := api.RetrieveReadingList(queries)
	if err != nil {
		return err
	}
	return nil
}

//ActivateSubcommand ...
func (c *ReadingListsCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
