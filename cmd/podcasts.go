package cmd

import (
	"fmt"
	"text/tabwriter"

	"github.com/Gealber/devto-cli/api"
)

type PodcastsCommand Command

func NewPodcastsCommand() *PodcastsCommand {
	return &PodcastsCommand{
		Name:        "podcasts",
		Description: "Retrieve podcasts",
		Subcommands: map[string]*Subcommand{
			"retrieve": {
				Description: "Retrieve podcasts",
				Active:      false,
			},
		},
	}
}

//Run execute the command
func (c *PodcastsCommand) Run() CommandValidationError {
	//Diferentiate two cases when is a retrieve and when is an update
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.Subcommands["retrieve"].Active {
		queries, err := processPodcastsQueries()
		if err != nil {
			return err
		}
		err = c.retrievePodcast(queries)
		if err != nil {
			return err
		}
	}
	return nil
}

//Validate check for the preconditions to execute this command
func (c *PodcastsCommand) Validate() CommandValidationError {
	//nothing to validate here
	return nil
}

//Helper display info about the command
func (c *PodcastsCommand) Helper(tw *tabwriter.Writer) {
	Helper(c.Name, c.Description, tw)
}

//SetData ...
func (c *PodcastsCommand) SetData(data string) {
	c.Data = data
}

//processPodcastsQueries read the data from the User input and put
//that data inside an ListingQuery structure
func processPodcastsQueries() (*api.PodcastEpisodesQuery, error) {
	//to store field from PodcastEpisodesQuery
	queries := new(api.PodcastEpisodesQuery)
	err := processInput(queries)
	if err != nil {
		return nil, err
	}
	return queries, nil
}

//retrievePodcast ...
func (c *PodcastsCommand) retrievePodcast(queries *api.PodcastEpisodesQuery) CommandValidationError {
	_, err := api.RetrievePodcastEpisodes(queries)
	if err != nil {
		return err
	}
	return nil
}

//ActivateSubcommand ...
func (c *PodcastsCommand) ActivateSubcommand(name string) error {
	if _, ok := c.Subcommands[name]; !ok {
		return NewCommandError(fmt.Sprintf("Subcommand %s not found", name))
	}
	c.Subcommands[name].Active = true
	return nil
}
